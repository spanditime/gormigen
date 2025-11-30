package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spanditime/gofin/tools/gormigen/internal/generate"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new migration",
	Long:  "Add a new migration package to the migrations directory",
	Args:  cobra.MinimumNArgs(1),
	Example: `
	gormigen add "Add users table"
	or
	gormigen add Add users table
	`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if len(args) > 1 {
			name = strings.Join(args, " ")
		} else if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
			return
		}
		err := generate.GenerateMigrationPackage(name)
		if err != nil {
			log.Fatalf("failed to add migration: %v", err)
			os.Exit(1)
			return
		}
		fmt.Println("migration added successfully")
	},
}
