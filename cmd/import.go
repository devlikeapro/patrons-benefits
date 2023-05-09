package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
import "github.com/devlikeapro/patrons-perks/internal/platforms"

var (
	filePath string
	platform string
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import patrons from a CSV file",
	Long:  `Import patrons from a CSV file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Importing patrons from file %s to platform %s\n", filePath, platform)
		err := platforms.ImportFromPlatform(platform, filePath)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringVarP(&filePath, "file", "f", "", "path to CSV file")
	importCmd.MarkFlagRequired("file")

	importCmd.Flags().StringVarP(&platform, "platform", "p", "", "platform to assign patrons to")
	importCmd.MarkFlagRequired("platform")
}
