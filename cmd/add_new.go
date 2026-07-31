package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"dumber/cmd/shared"

	"github.com/spf13/cobra"
)

func findExisting(file []shared.AddURIStruct, name, uri string) *shared.AddURIStruct {
	for i := range file {
		if strings.EqualFold(file[i].Name, name) || strings.EqualFold(file[i].Uri, uri) {
			return &file[i]
		}
	}
	return nil
}

var scanner = bufio.NewScanner(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func CobraCreate() *cobra.Command {
	return &cobra.Command{
		Use: "add",
		RunE: func(cmd *cobra.Command, args []string) error {

			name := readLine("Please enter name of " + "... ")
			db := readLine("Please enter db of " + name + "... ")
			uri := readLine("Please add uri of " + name + "... ")
			status := readLine("Please enter status of " + name + "... ")

			file, err := shared.LoadExistingData()
			if err != nil {
				return err
			}

			if existing := findExisting(file, name, uri); existing != nil {
				fmt.Println("data already exists:")
				fmt.Printf("%+v\n", *existing)
				return nil
			}

			file = append(file, shared.AddURIStruct{
				ID:     int32(len(file) + 1),
				Uri:    uri,
				DB:     db,
				Name:   name,
				Status: status,
			})

			jsonFile, err := json.MarshalIndent(file, "", "  ")

			if err != nil {
				return err
			}

			os.Mkdir("json", 0755)
			os.WriteFile(shared.DataFilePath, jsonFile, 0644)
			fmt.Println("data added successfully")
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(CobraCreate())
}
