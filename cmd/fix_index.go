package cmd

import (
	"fmt"
	"log"

	"github.com/spanditime/gofin/tools/gormigen/pkg/generate"
	"github.com/spf13/cobra"
)

var fixIndexCmd = &cobra.Command{
	Use:   "fix-index [strategy:auto|inplace|push-back|dry-run]",
	Short: "Fix index collisions",
	Long:  "Fix index collisions in migrations, by renaming migrations with collisions to a new version. Default strategy is auto.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		strategy := generate.ModeAuto
		if len(args) > 0 {
			strategy = generate.Mode(args[0])
			if strategy != generate.ModeAuto &&
				strategy != generate.ModeInplace &&
				strategy != generate.ModePushBack &&
				strategy != generate.ModeDryRun {
				log.Fatalf("invalid strategy: %s", args[0])
				return
			}
		}
		err := generate.FixCollisions(strategy)
		if err != nil {
			log.Fatalf("failed to fix collisions: %v", err)
		}
		fmt.Println("collisions fixed successfully")
	},
}
