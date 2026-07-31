package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type AddURIStruct struct {
	ID     int32  `json:"id"`
	DB     string `json:"db"`
	Uri    string `json:"uri"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

const dataFilePath string = "json/data.json"

func loadExistingData() ([]AddURIStruct, error) {
	file := []AddURIStruct{}

	raw, err := os.ReadFile(dataFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return file, nil
		}
		return nil, err
	}

	if len(strings.TrimSpace(string(raw))) == 0 {
		return file, nil
	}

	if err := json.Unmarshal(raw, &file); err != nil {
		return nil, err
	}

	return file, nil
}

func findExisting(file []AddURIStruct, name, uri string) *AddURIStruct {
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

			file, err := loadExistingData()
			if err != nil {
				return err
			}

			if existing := findExisting(file, name, uri); existing != nil {
				fmt.Println("data already exists:")
				fmt.Printf("%+v\n", *existing)
				return nil
			}

			file = append(file, AddURIStruct{
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
			os.WriteFile(dataFilePath, jsonFile, 0644)
			fmt.Println("data added successfully")
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(CobraCreate())
}
