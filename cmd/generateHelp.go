package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gHelpCmd represents the generateHelp command
var gHelpCmd = &cobra.Command{
	Use:     "help",
	Short:   "[h] Generate a help file for all the existing command in html format.",
	Long:    `[h] Generate a help file for all the existing command in html format.`,
	Aliases: []string{"h"},
	Example: `
  app generate help
  app g h
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("\n Generating help file in HTML foramt. \n\n")
		return nil
	},
}

func init() {
	generateCmd.AddCommand(gHelpCmd)
}
