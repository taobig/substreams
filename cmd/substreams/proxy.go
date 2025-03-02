package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"connectrpc.com/connect"
	grpcreflect "connectrpc.com/grpcreflect"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/streamingfast/cli/sflags"
	"github.com/streamingfast/substreams/client"
	"github.com/streamingfast/substreams/manifest"
	pbrpcsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/rpc/v2"
	ssconnect "github.com/streamingfast/substreams/pb/sf/substreams/rpc/v2/pbsubstreamsrpcconnect"
	"github.com/streamingfast/substreams/tools"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc/metadata"
)

var proxyCmd = &cobra.Command{
	Use:          "proxy <package>",
	Short:        "A tool to proxy substreams requests from a web browser (connect-web protocol)",
	RunE:         runProxy,
	SilenceUsage: true,
}

type ConnectServer struct {
	ssconnect.UnimplementedStreamHandler
	Manifest               string
	StartBlock             uint64
	SubstreamsClientConfig *client.SubstreamsClientConfig
}

func (cs *ConnectServer) Blocks(
	ctx context.Context,
	req *connect.Request[pbrpcsubstreams.Request],
	stream *connect.ServerStream[pbrpcsubstreams.Response],
) error {
	newReq := &pbrpcsubstreams.Request{
		StartBlockNum:                       req.Msg.StartBlockNum,
		StopBlockNum:                        req.Msg.StopBlockNum,
		StartCursor:                         req.Msg.StartCursor,
		FinalBlocksOnly:                     req.Msg.FinalBlocksOnly,
		OutputModule:                        req.Msg.OutputModule,
		Modules:                             req.Msg.Modules,
		DebugInitialStoreSnapshotForModules: req.Msg.DebugInitialStoreSnapshotForModules,
	}

	if cs.Manifest != "" {
		manifestReader, err := manifest.NewReader(cs.Manifest)
		if err != nil {
			return fmt.Errorf("manifest reader: %w", err)
		}

		pkg, _, err := manifestReader.Read()
		if err != nil {
			return fmt.Errorf("read manifest %q: %w", cs.Manifest, err)
		}
		newReq.Modules = pkg.Modules
	}

	if cs.StartBlock != 0 {
		newReq.StartBlockNum = int64(cs.StartBlock)
	}

	ssClient, connClose, callOpts, headers, err := client.NewSubstreamsClient(cs.SubstreamsClientConfig)
	if err != nil {
		return fmt.Errorf("substreams client setup: %w", err)
	}
	defer connClose()

	if err := newReq.Validate(); err != nil {
		return fmt.Errorf("validate request: %w", err)
	}

	if headers.IsSet() {
		ctx = metadata.AppendToOutgoingContext(ctx, headers.ToArray()...)
	}
	cli, err := ssClient.Blocks(ctx, newReq, callOpts...)
	if err != nil {
		return fmt.Errorf("call sf.substreams.rpc.v2.Stream/Blocks: %w", err)
	}

	for {
		resp, err := cli.Recv()
		if resp != nil {
			stream.Send(resp)
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func init() {
	proxyCmd.Flags().StringP("substreams-endpoint", "e", "mainnet.eth.streamingfast.io:443", "Substreams gRPC endpoint")
	proxyCmd.Flags().String("substreams-api-token-envvar", "SUBSTREAMS_API_TOKEN", "name of variable containing Substreams Authentication token")
	proxyCmd.Flags().String("substreams-api-key-envvar", "SUBSTREAMS_API_KEY", "Name of variable containing Substreams Api Key")
	proxyCmd.Flags().String("listen-addr", "localhost:8080", "listen on this address (unencrypted)")
	proxyCmd.Flags().BoolP("insecure", "k", false, "Skip certificate validation on GRPC connection")
	proxyCmd.Flags().BoolP("plaintext", "p", false, "Establish GRPC connection in plaintext")
	proxyCmd.Flags().String("force-manifest", "", "if non-empty, the requests' modules will be replaced by the modules loaded from this location. Can be a local spkg or yaml file, or a remote (HTTP) spkg file.")
	proxyCmd.Flags().Uint64("force-start-block", 0, "if non-zero, the requests' start-block will be replaced by this value")

	tools.Cmd.AddCommand(proxyCmd)
}

func runProxy(cmd *cobra.Command, args []string) error {
	addr := sflags.MustGetString(cmd, "listen-addr")
	fmt.Println("listening on", addr)

	authToken, authType := tools.GetAuth(cmd, "substreams-api-key-envvar", "substreams-api-token-envvar")
	substreamsClientConfig := client.NewSubstreamsClientConfig(
		sflags.MustGetString(cmd, "substreams-endpoint"),
		authToken,
		authType,
		sflags.MustGetBool(cmd, "insecure"),
		sflags.MustGetBool(cmd, "plaintext"),
	)

	cs := &ConnectServer{
		Manifest:               sflags.MustGetString(cmd, "force-manifest"),
		SubstreamsClientConfig: substreamsClientConfig,
		StartBlock:             sflags.MustGetUint64(cmd, "force-start-block"),
	}

	reflector := grpcreflect.NewStaticReflector(
		"sf.substreams.rpc.v2.Stream",
	)

	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.
	mux.Handle(ssconnect.NewStreamHandler(cs))
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return http.ListenAndServe(
		addr,
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(
			newCORS().Handler(mux),
			&http2.Server{}),
	)
}

func newCORS() *cors.Cors {
	// To let web developers play with the demo service from browsers, we need a
	// very permissive CORS setup.
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
}
