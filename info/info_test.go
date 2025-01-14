package info

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/streamingfast/substreams/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBasicInfo(t *testing.T) {
	reader, err := manifest.NewReader("https://github.com/streamingfast/substreams-uniswap-v3/releases/download/v0.2.8/substreams.spkg")
	require.NoError(t, err)

	pkg, graph, err := reader.Read()
	require.NoError(t, err)

	info, err := Basic(pkg, graph)
	require.NoError(t, err)

	r, err := json.MarshalIndent(info, "", "  ")
	require.NoError(t, err)

	fmt.Println(string(r))
}

func TestExtendedInfo(t *testing.T) {
	info, err := Extended("https://github.com/streamingfast/substreams-uniswap-v3/releases/download/v0.2.8/substreams.spkg", "graph_out", false, 0)
	require.NoError(t, err)

	r, err := json.MarshalIndent(info, "", "  ")
	require.NoError(t, err)

	assert.Equal(t, uint64(12369621), info.Modules[0].InitialBlock)

	fmt.Println(string(r))
}

func TestExtendedInfoFirstStreamable(t *testing.T) {
	info, err := Extended("https://github.com/streamingfast/substreams-uniswap-v3/releases/download/v0.2.8/substreams.spkg", "graph_out", false, 999999999)
	require.NoError(t, err)

	assert.Equal(t, uint64(999999999), info.Modules[0].InitialBlock)
	assert.Equal(t, uint64(999999999), info.Modules[1].InitialBlock)
	assert.Equal(t, uint64(999999999), info.Modules[2].InitialBlock)
	// ...
}
