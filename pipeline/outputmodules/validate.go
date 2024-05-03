package outputmodules

import (
	"fmt"

	"github.com/streamingfast/substreams/manifest"
	pbssinternal "github.com/streamingfast/substreams/pb/sf/substreams/intern/v2"
	pbsubstreamsrpc "github.com/streamingfast/substreams/pb/sf/substreams/rpc/v2"
	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
)

// Deprecated: use ValidateTier1Request
var ValidateRequest = ValidateTier1Request

// ValidateTier1Request is run by the server code.
func ValidateTier1Request(request *pbsubstreamsrpc.Request, blockType string) error {
	if err := request.Validate(); err != nil {
		return fmt.Errorf("validate tier1 request: %s", err)
	}

	err := validateRequest(request.Modules.Binaries, request.Modules, request.OutputModule, blockType)
	if err != nil {
		return err
	}

	return nil
}

func ValidateTier2Request(request *pbssinternal.ProcessRangeRequest) error {
	if err := request.Validate(); err != nil {
		return fmt.Errorf("validate tier2 request: %s", err)
	}

	err := validateRequest(request.Modules.Binaries, request.Modules, request.OutputModule, request.BlockType)
	if err != nil {
		return err
	}

	return nil
}

func validateRequest(binaries []*pbsubstreams.Binary, modules *pbsubstreams.Modules, outputModule string, blockType string) error {
	if err := validateBinaryTypes(binaries); err != nil {
		return err
	}

	if err := manifest.ValidateModules(modules); err != nil {
		return fmt.Errorf("modules validation failed: %w", err)
	}

	if err := validateModuleGraph(modules.Modules, outputModule, blockType); err != nil {
		return err
	}

	if err := checkNotImplemented(modules.Modules); err != nil {
		return fmt.Errorf("checking feature not implemented: %w", err)
	}
	return nil
}

func checkNotImplemented(mods []*pbsubstreams.Module) error {
	for _, mod := range mods {
		if mod.ModuleKind() == pbsubstreams.ModuleKindBlockIndex {
			return fmt.Errorf("block index module is not implemented")
		}
		if mod.GetBlockFilter() != nil {
			return fmt.Errorf("block filter module is not implemented")
		}
	}
	return nil
}

func validateModuleGraph(mods []*pbsubstreams.Module, outputModule string, blockType string) error {
	graph, err := manifest.NewModuleGraph(mods)
	if err != nil {
		return fmt.Errorf("should have been able to derive modules graph: %w", err)
	}

	// Already validated by `ValidateTier1Request` above, so we can use the `Must...` version
	ancestors, err := graph.AncestorsOf(outputModule)
	if err != nil {
		return fmt.Errorf("computing ancestors of %q: %w", outputModule, err)
	}

	// We must only validate the input source against module that we are going to actually run. A Substreams
	// could provide modules for multiple chain while executing only one of them in which case only the one
	// run (and its dependencies transitively) should be checked.
	for _, mod := range ancestors {
		for _, input := range mod.Inputs {
			if src := input.GetSource(); src != nil {
				if src.Type != blockType && src.Type != "sf.substreams.v1.Clock" {
					return fmt.Errorf("input source %q not supported, only %q and 'sf.substreams.v1.Clock' are valid", src, blockType)
				}
			}
		}
	}

	return nil
}

func validateBinaryTypes(bins []*pbsubstreams.Binary) error {
	for _, binary := range bins {
		switch binary.Type {
		case "wasm/rust-v1":
		case "wasip1/tinygo-v1":
		default:
			return fmt.Errorf(`unsupported binary type: %q, please use "wasm/rust-v1" or "wasip1/tinygo-v1"`, binary.Type)
		}
	}
	return nil
}
