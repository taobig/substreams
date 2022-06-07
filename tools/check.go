package tools

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/streamingfast/dstore"
	"github.com/streamingfast/substreams/state"
)

var checkCmd = &cobra.Command{
	Use:   "check <store_url>",
	Short: "checks the integrity of the kv files in a given store",
	Args:  cobra.ExactArgs(1),
	RunE:  checkE,
}

func init() {
	Cmd.AddCommand(checkCmd)
}

func checkE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	store, err := dstore.NewStore(args[0], "", "", false)
	if err != nil {
		return fmt.Errorf("could not create store from %s: %w", args[0], err)
	}

	builder := state.Store{
		Store: store,
	}

	snapshots, err := builder.ListSnapshots(ctx)
	if err != nil {
		return fmt.Errorf("listing snapshots: %w", err)
	}

	sort.Slice(snapshots.Files, func(i, j int) bool {
		return snapshots.Files[i].Range.ExclusiveEndBlock < snapshots.Files[j].Range.ExclusiveEndBlock
	})

	var prevSnapshot *state.Snapshot
	for _, snapshot := range snapshots.Files {
		if prevSnapshot == nil {
			prevSnapshot = &snapshot
			continue
		}

		if snapshot.StartBlock != prevSnapshot.ExclusiveEndBlock {
			return fmt.Errorf("**hole found** between %d and %d", prevSnapshot.Range.ExclusiveEndBlock, snapshot.Range.ExclusiveEndBlock)
		}

		prevSnapshot = &snapshot
	}

	return err
}
