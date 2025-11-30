package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spanditime/gofin/tools/gormigen/internal/config"
	"github.com/spanditime/gofin/tools/gormigen/internal/generate"
	"github.com/spf13/cobra"
)

var ignoreNewer bool

var addCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Add a new migration",
	Long: `Add a new migration package to the migrations directory
To ignore checking for newer migrations, use --ignore-newer flag`,
	Args: cobra.ExactArgs(1),
	Example: `
gormigen add "Add users table"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		cfg, err := config.GetConfig()
		if err != nil {
			log.Fatalf("failed to get config: %v", err)
			os.Exit(1)
			return
		}
		err = generate.GenerateMigrationPackage(cfg.Migrations, time.Now(), name, !ignoreNewer)
		if err != nil {
			log.Fatalf("failed to add migration: %v", err)
			os.Exit(1)
			return
		}
		fmt.Println("migration added successfully")
	},
}

func init() {
	addCmd.Flags().BoolVarP(&ignoreNewer, "ignore-newer", "i", false, "ignore checking for newer migrations")
	rootCmd.AddCommand(addCmd)
}
