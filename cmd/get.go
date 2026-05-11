package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func GetAllJson() *cobra.Command {
	first_flag := &cobra.Command{
		Use:   "list",
		Short: "Get json saved data",
		RunE: func(cmd *cobra.Command, args []string) error {
			nameFiles, err := os.ReadDir("json")

			if err != nil {
				return err
			}

			for _, data := range nameFiles {
				fmt.Printf("name of file: %s\n", data.Name())
			}
			return nil
		},
	}

	return first_flag
}

func GetJsonData() *cobra.Command {
	first_flag := &cobra.Command{
		Use:   "get",
		Short: "Get json saved data",
	}

	last_flag := &cobra.Command{
		Use:   "open",
		Short: "Get All json data",
		RunE: func(cmd *cobra.Command, args []string) error {
			nameFile := args[0]

			file, err := os.ReadFile("json/" + nameFile + ".json")
			if err != nil {
				return err
			}

			var data []AddURIStruct
			err = json.Unmarshal(file, &data)
			if err != nil {
				return err
			}

			for _, data := range data {
				fmt.Printf("Name of file: %s\n", data.Name)
				fmt.Printf("Name of Database: %s\n", data.DB)
				fmt.Printf("Database uri: %s\n", data.Uri)
				fmt.Printf("Status of Database: %s\n", data.Status)
			}
			return nil
		},
	}
	first_flag.AddCommand(last_flag)

	return first_flag
}

func init() {
	rootCmd.AddCommand(GetAllJson())
	rootCmd.AddCommand(GetJsonData())
}
