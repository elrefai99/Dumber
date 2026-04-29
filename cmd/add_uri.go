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
	NameFile string `json:"name_file"`
	Uri      string `json:"uri"`
	Status   string `json:"status"`
	Name     string `json:"name"`
}

var scanner = bufio.NewScanner(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func CreateAsDefault() {

}

func CobraCreate() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Add new uri",
		RunE: func(cmd *cobra.Command, args []string) error {

			name := readLine("Please enter name of " + "... ")
			nameFile := readLine("Please enter name of file " + "... ")
			uri := readLine("Please add uri of " + name + "... ")
			status := readLine("Please enter status of " + name + "... ")

			file := []AddURIStruct{}
			file = append(file, AddURIStruct{
				Uri:      uri,
				Name:     name,
				Status:   status,
				NameFile: nameFile,
			})

			jsonFile, err := json.MarshalIndent(file, "", "  ")

			if err != nil {
				return err
			}

			os.Mkdir("json", 0755)
			os.WriteFile("json/"+nameFile+".json", jsonFile, 0644)
			fmt.Println("data added successfully")
			return nil
		},
	}
}
func init() {
	rootCmd.AddCommand(CobraCreate())
}
