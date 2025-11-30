package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gormigen",
	Short: "Gormigen is a tool for generating boilerplate for gormigrations",
	Long: `Gormigen is a tool for generating boilerplate for gormigrations,
	that helps you easily create, manage and apply migrations to your database`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(fixIndexCmd)
}
