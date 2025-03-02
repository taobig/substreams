package index

import (
	"fmt"

	"github.com/streamingfast/dstore"
	"github.com/streamingfast/substreams/manifest"
	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
	"go.uber.org/zap"
)

type Configs struct {
	ConfigMap map[string]*Config
	logger    *zap.Logger
}

func NewConfigs(baseObjectStore dstore.Store, allRequestedModules []*pbsubstreams.Module, moduleHashes *manifest.ModuleHashes, logger *zap.Logger) (*Configs, error) {
	out := make(map[string]*Config)
	for _, mod := range allRequestedModules {
		conf, err := NewConfig(
			mod.Name,
			mod.InitialBlock,
			moduleHashes.Get(mod.Name),
			baseObjectStore,
			logger,
		)
		if err != nil {
			return nil, fmt.Errorf("new index config for %q: %w", mod.Name, err)
		}
		out[mod.Name] = conf
	}

	return &Configs{
		ConfigMap: out,
		logger:    logger,
	}, nil
}
