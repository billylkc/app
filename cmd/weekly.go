package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var wCmd = &cobra.Command{
	Use:     "weekly [-d date] [previous] [nrecords]",
	Short:   "[w] Collection of weekly queries",
	Long:    `[w] Collection of weekly queries`,
	Aliases: []string{"w"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(wCmd)
}
