package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var dailyCmd = &cobra.Command{
	Use:     "daily",
	Short:   "Collection of daily query",
	Long:    `Collection of daily query`,
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(dailyCmd)
}
