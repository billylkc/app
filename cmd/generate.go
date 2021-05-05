package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "[g] Experimental. Generate help commands, reports, etc.",
	Long:    `[g] Experimental. Generate help commands, reports, etc.`,
	Aliases: []string{"g"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
