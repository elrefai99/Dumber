package cmd

import (
	"dumber/cmd/shared"
	"fmt"

	"github.com/spf13/cobra"
)

func CobraGet() *cobra.Command {
	return &cobra.Command{
		Use: "get",
		RunE: func(cmd *cobra.Command, args []string) error {
			file, err := shared.LoadExistingData()

			if err != nil {
				return err
			}

			fmt.Print(file)
			return err
		},
	}
}

func init() {
	rootCmd.AddCommand(CobraGet())
}
