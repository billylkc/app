package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var qCmd = &cobra.Command{
	Use:     "query",
	Short:   "[q] Detailed query of products, members, etc",
	Long:    `[q] Detailed query of products, members, etc`,
	Aliases: []string{"q"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(qCmd)
}
