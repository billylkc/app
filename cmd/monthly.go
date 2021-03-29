package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var mCmd = &cobra.Command{
	Use:     "monthly",
	Short:   "Collection of monthly query",
	Long:    `Collection of monthly query`,
	Aliases: []string{"m"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(mCmd)
}
