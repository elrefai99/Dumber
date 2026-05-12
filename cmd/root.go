package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "This CLI tool is designed to simplify the process of exporting MongoDB data to JSON files.\n It allows you to save your database structures and content in a format that is easy to backup, share, and restore.",
	Long:  "This CLI tool is designed to simplify the process of exporting MongoDB data to JSON files.\n It allows you to save your database structures and content in a format that is easy to backup, share, and restore.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
