package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"dumber/internal/dumb"

	"github.com/spf13/cobra"
)

func DumbDB() *cobra.Command {

	database := &cobra.Command{
		Use:   "dumb",
		Short: "Export database",
		RunE: func(cmd *cobra.Command, args []string) error {
			fileName := args[0]

			fileData, err := os.ReadFile("json/" + fileName + ".json")

			if err != nil {
				return err
			}

			var dbSettings []AddURIStruct
			err = json.Unmarshal(fileData, &dbSettings)
			if err != nil {
				return err
			}

			os.Mkdir("backup", 0755)
			for _, data := range dbSettings {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					return err
				}
				out := filepath.Join(homeDir, "Downloads", "backup", "dumber-export", time.Now().Format("2006-01-02"))
				dumb.Dumb(data.Name, out, data.Uri, true)

			}
			// db name
			return nil
		},
	}

	return database
}

func init() {
	rootCmd.AddCommand(DumbDB())
}
