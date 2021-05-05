package cmd

import (
	"fmt"

	"github.com/billylkc/app/app"
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
		commands := app.GetCommandStruct()
		fmt.Println(PrettyPrint(commands))

		// TODO: Print to html

		return nil
	},
}

func init() {
	generateCmd.AddCommand(gHelpCmd)
}
