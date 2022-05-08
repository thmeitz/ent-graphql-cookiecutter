package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"go.uber.org/zap"
)

// genDocsCmd represents the genDocs command
var genDocsCmd = &cobra.Command{
	Use:   "gen-md-docs",
	Short: "generates cli doc",
	Run: func(cmd *cobra.Command, args []string) {
		logger := zap.NewExample()

		if len(args) != 1 {
			if err := cmd.Usage(); err != nil {
				logger.Fatal("Failed to print usage", zap.Error(err))
			}
			fmt.Println()
			logger.Fatal(`Missing output path`)
		}

		path := args[0]
		if path[len(path)-1] != os.PathSeparator {
			path += string(os.PathSeparator)
		}
		if err := doc.GenMarkdownTree(rootCmd, path); err != nil {
			logger.Fatal("Could no generate markdown",
				zap.Any(`filename`, args[0]),
				zap.Error(err),
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(genDocsCmd)
}
