package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var dCmd = &cobra.Command{
	Use:     "daily",
	Short:   "[d] Collection of daily queries",
	Long:    `[d] Collection of daily queries`,
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(dCmd)
}
