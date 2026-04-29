package cmd

import (
	"github.com/spf13/cobra"
)

func DumbDB() *cobra.Command {
	database := &cobra.Command{
		Use:   "dumb",
		Short: "Dumb database",
	}

	export := &cobra.Command{
		Use:   "export",
		Short: "Export database",
		RunE: func(cmd *cobra.Command, args []string) error {
			// db name
			Dumb()
			return nil
		},
	}

	export.AddCommand(database)

	return export
}

func init() {
	rootCmd.AddCommand(DumbDB())
}
