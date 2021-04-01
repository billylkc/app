package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// topCmd represents the top command
var tCmd = &cobra.Command{
	Use:     "top",
	Short:   "Top sales for members, products, etc..",
	Long:    `Top sales for members, products, etc..`,
	Aliases: []string{"t"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tCmd)

}
