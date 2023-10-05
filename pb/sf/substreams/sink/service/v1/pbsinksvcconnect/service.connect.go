// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: sf/substreams/sink/service/v1/service.proto

package pbsinksvcconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/streamingfast/substreams/pb/sf/substreams/sink/service/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ProviderName is the fully-qualified name of the Provider service.
	ProviderName = "sf.substreams.sink.service.v1.Provider"
)

// ProviderClient is a client for the sf.substreams.sink.service.v1.Provider service.
type ProviderClient interface {
	Deploy(context.Context, *connect_go.Request[v1.DeployRequest]) (*connect_go.Response[v1.DeployResponse], error)
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	Info(context.Context, *connect_go.Request[v1.InfoRequest]) (*connect_go.Response[v1.InfoResponse], error)
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
	Pause(context.Context, *connect_go.Request[v1.PauseRequest]) (*connect_go.Response[v1.PauseResponse], error)
	Stop(context.Context, *connect_go.Request[v1.StopRequest]) (*connect_go.Response[v1.StopResponse], error)
	Resume(context.Context, *connect_go.Request[v1.ResumeRequest]) (*connect_go.Response[v1.ResumeResponse], error)
	Remove(context.Context, *connect_go.Request[v1.RemoveRequest]) (*connect_go.Response[v1.RemoveResponse], error)
}

// NewProviderClient constructs a client for the sf.substreams.sink.service.v1.Provider service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProviderClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ProviderClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &providerClient{
		deploy: connect_go.NewClient[v1.DeployRequest, v1.DeployResponse](
			httpClient,
			baseURL+"/sf.substreams.sink.service.v1.Provider/Deploy",
			opts...,
		),
		update: connect_go.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+"/sf.substreams.sink.service.v1.Provider/Update",
			opts...,
		),
		info: connect_go.NewClient[v1.InfoRequest, v1.InfoResponse](
			httpClient,
			baseURL+"/sf.substreams.sink.service.v1.Provider/Info",
			opts...,
		),
		list: connect_go.NewClient[v1.ListRequest, v1.ListResponse](
			httpClient,
			baseURL+"/sf.substreams.sink.service.v1.Provider/List",
			opts...,
		),
		pause: connect_go.NewClient[v1.PauseRequest, v1.PauseResponse](
			httpClient,
			baseURL+"/sf.substreams.sink.service.v1.Provider/Pause",
			opts...,
		),
		stop: connect_go.NewClient[v1.StopRequest, v1.StopResponse](
			httpClient,
			baseURL+"/sf.substreams.sink.service.v1.Provider/Stop",
			opts...,
		),
		resume: connect_go.NewClient[v1.ResumeRequest, v1.ResumeResponse](
			httpClient,
			baseURL+"/sf.substreams.sink.service.v1.Provider/Resume",
			opts...,
		),
		remove: connect_go.NewClient[v1.RemoveRequest, v1.RemoveResponse](
			httpClient,
			baseURL+"/sf.substreams.sink.service.v1.Provider/Remove",
			opts...,
		),
	}
}

// providerClient implements ProviderClient.
type providerClient struct {
	deploy *connect_go.Client[v1.DeployRequest, v1.DeployResponse]
	update *connect_go.Client[v1.UpdateRequest, v1.UpdateResponse]
	info   *connect_go.Client[v1.InfoRequest, v1.InfoResponse]
	list   *connect_go.Client[v1.ListRequest, v1.ListResponse]
	pause  *connect_go.Client[v1.PauseRequest, v1.PauseResponse]
	stop   *connect_go.Client[v1.StopRequest, v1.StopResponse]
	resume *connect_go.Client[v1.ResumeRequest, v1.ResumeResponse]
	remove *connect_go.Client[v1.RemoveRequest, v1.RemoveResponse]
}

// Deploy calls sf.substreams.sink.service.v1.Provider.Deploy.
func (c *providerClient) Deploy(ctx context.Context, req *connect_go.Request[v1.DeployRequest]) (*connect_go.Response[v1.DeployResponse], error) {
	return c.deploy.CallUnary(ctx, req)
}

// Update calls sf.substreams.sink.service.v1.Provider.Update.
func (c *providerClient) Update(ctx context.Context, req *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Info calls sf.substreams.sink.service.v1.Provider.Info.
func (c *providerClient) Info(ctx context.Context, req *connect_go.Request[v1.InfoRequest]) (*connect_go.Response[v1.InfoResponse], error) {
	return c.info.CallUnary(ctx, req)
}

// List calls sf.substreams.sink.service.v1.Provider.List.
func (c *providerClient) List(ctx context.Context, req *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Pause calls sf.substreams.sink.service.v1.Provider.Pause.
func (c *providerClient) Pause(ctx context.Context, req *connect_go.Request[v1.PauseRequest]) (*connect_go.Response[v1.PauseResponse], error) {
	return c.pause.CallUnary(ctx, req)
}

// Stop calls sf.substreams.sink.service.v1.Provider.Stop.
func (c *providerClient) Stop(ctx context.Context, req *connect_go.Request[v1.StopRequest]) (*connect_go.Response[v1.StopResponse], error) {
	return c.stop.CallUnary(ctx, req)
}

// Resume calls sf.substreams.sink.service.v1.Provider.Resume.
func (c *providerClient) Resume(ctx context.Context, req *connect_go.Request[v1.ResumeRequest]) (*connect_go.Response[v1.ResumeResponse], error) {
	return c.resume.CallUnary(ctx, req)
}

// Remove calls sf.substreams.sink.service.v1.Provider.Remove.
func (c *providerClient) Remove(ctx context.Context, req *connect_go.Request[v1.RemoveRequest]) (*connect_go.Response[v1.RemoveResponse], error) {
	return c.remove.CallUnary(ctx, req)
}

// ProviderHandler is an implementation of the sf.substreams.sink.service.v1.Provider service.
type ProviderHandler interface {
	Deploy(context.Context, *connect_go.Request[v1.DeployRequest]) (*connect_go.Response[v1.DeployResponse], error)
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	Info(context.Context, *connect_go.Request[v1.InfoRequest]) (*connect_go.Response[v1.InfoResponse], error)
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
	Pause(context.Context, *connect_go.Request[v1.PauseRequest]) (*connect_go.Response[v1.PauseResponse], error)
	Stop(context.Context, *connect_go.Request[v1.StopRequest]) (*connect_go.Response[v1.StopResponse], error)
	Resume(context.Context, *connect_go.Request[v1.ResumeRequest]) (*connect_go.Response[v1.ResumeResponse], error)
	Remove(context.Context, *connect_go.Request[v1.RemoveRequest]) (*connect_go.Response[v1.RemoveResponse], error)
}

// NewProviderHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProviderHandler(svc ProviderHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/sf.substreams.sink.service.v1.Provider/Deploy", connect_go.NewUnaryHandler(
		"/sf.substreams.sink.service.v1.Provider/Deploy",
		svc.Deploy,
		opts...,
	))
	mux.Handle("/sf.substreams.sink.service.v1.Provider/Update", connect_go.NewUnaryHandler(
		"/sf.substreams.sink.service.v1.Provider/Update",
		svc.Update,
		opts...,
	))
	mux.Handle("/sf.substreams.sink.service.v1.Provider/Info", connect_go.NewUnaryHandler(
		"/sf.substreams.sink.service.v1.Provider/Info",
		svc.Info,
		opts...,
	))
	mux.Handle("/sf.substreams.sink.service.v1.Provider/List", connect_go.NewUnaryHandler(
		"/sf.substreams.sink.service.v1.Provider/List",
		svc.List,
		opts...,
	))
	mux.Handle("/sf.substreams.sink.service.v1.Provider/Pause", connect_go.NewUnaryHandler(
		"/sf.substreams.sink.service.v1.Provider/Pause",
		svc.Pause,
		opts...,
	))
	mux.Handle("/sf.substreams.sink.service.v1.Provider/Stop", connect_go.NewUnaryHandler(
		"/sf.substreams.sink.service.v1.Provider/Stop",
		svc.Stop,
		opts...,
	))
	mux.Handle("/sf.substreams.sink.service.v1.Provider/Resume", connect_go.NewUnaryHandler(
		"/sf.substreams.sink.service.v1.Provider/Resume",
		svc.Resume,
		opts...,
	))
	mux.Handle("/sf.substreams.sink.service.v1.Provider/Remove", connect_go.NewUnaryHandler(
		"/sf.substreams.sink.service.v1.Provider/Remove",
		svc.Remove,
		opts...,
	))
	return "/sf.substreams.sink.service.v1.Provider/", mux
}

// UnimplementedProviderHandler returns CodeUnimplemented from all methods.
type UnimplementedProviderHandler struct{}

func (UnimplementedProviderHandler) Deploy(context.Context, *connect_go.Request[v1.DeployRequest]) (*connect_go.Response[v1.DeployResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sf.substreams.sink.service.v1.Provider.Deploy is not implemented"))
}

func (UnimplementedProviderHandler) Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sf.substreams.sink.service.v1.Provider.Update is not implemented"))
}

func (UnimplementedProviderHandler) Info(context.Context, *connect_go.Request[v1.InfoRequest]) (*connect_go.Response[v1.InfoResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sf.substreams.sink.service.v1.Provider.Info is not implemented"))
}

func (UnimplementedProviderHandler) List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sf.substreams.sink.service.v1.Provider.List is not implemented"))
}

func (UnimplementedProviderHandler) Pause(context.Context, *connect_go.Request[v1.PauseRequest]) (*connect_go.Response[v1.PauseResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sf.substreams.sink.service.v1.Provider.Pause is not implemented"))
}

func (UnimplementedProviderHandler) Stop(context.Context, *connect_go.Request[v1.StopRequest]) (*connect_go.Response[v1.StopResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sf.substreams.sink.service.v1.Provider.Stop is not implemented"))
}

func (UnimplementedProviderHandler) Resume(context.Context, *connect_go.Request[v1.ResumeRequest]) (*connect_go.Response[v1.ResumeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sf.substreams.sink.service.v1.Provider.Resume is not implemented"))
}

func (UnimplementedProviderHandler) Remove(context.Context, *connect_go.Request[v1.RemoveRequest]) (*connect_go.Response[v1.RemoveResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sf.substreams.sink.service.v1.Provider.Remove is not implemented"))
}
