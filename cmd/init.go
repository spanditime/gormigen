package cmd

import (
	"log"

	"github.com/spanditime/gofin/tools/gormigen/pkg/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize gormigen config file",
	Long:  "Initialize gormigen",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
