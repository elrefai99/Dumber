package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UpdateUri() *cobra.Command {

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update uri data",
	}
	updateCommand := &cobra.Command{
		Use:   "uri",
		Short: "Update uri data",
		RunE: func(cmd *cobra.Command, args []string) error {
			nameFile := readLine("Please enter name of file " + "... ")
			name := readLine("Please enter name of " + "... ")
			db := readLine("Please enter db of " + name + "... ")
			uri := readLine("Please add uri of " + name + "... ")
			status := readLine("Please enter status of " + name + "... ")

			file, err := os.ReadFile("json/" + nameFile + ".json")
			if err != nil {
				return err
			}
			var data []AddURIStruct
			err = json.Unmarshal(file, &data)

			if err != nil {
				return err
			}
			for i := range data {
				if name != "" {
					data[i].Name = name
				}
				if db != "" {
					data[i].DB = db
				}
				if uri != "" {
					data[i].Uri = uri
				}
				if status != "" {
					data[i].Status = status
				}
			}
			jsonFile, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return err
			}
			os.WriteFile("json/"+nameFile+".json", jsonFile, 0644)
			fmt.Println("data updated successfully")
			return nil
		},
	}

	updateCmd.AddCommand(updateCommand)

	return updateCmd
}

func init() {
	rootCmd.AddCommand(UpdateUri())
}
