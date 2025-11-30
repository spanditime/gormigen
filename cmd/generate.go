package cmd

import (
	"log"

	"github.com/spanditime/gofin/tools/gormigen/internal/config"
	"github.com/spanditime/gofin/tools/gormigen/internal/generate"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate boilerplate for gormigrations",
	Long:  `Generate boilerplate for gormigrations, that helps you easily create, manage and apply migrations to your database`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.GetConfig()
		if err != nil {
			log.Fatalf("failed to get config: %v", err)
		}
		err = generate.Generate(cfg)
		if err != nil {
			log.Fatalf("failed to generate: %v", err)
		}
	},
}
