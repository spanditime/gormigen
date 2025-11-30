package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize gormigen config file [WIP]",
	Long:  "Initialize gormigen command is not implemented yet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init is not implemented yet")
		os.Exit(1)
	},
}
