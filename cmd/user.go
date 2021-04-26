package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:     "user",
	Short:   "[u] Connection of queries about user.",
	Long:    `[u] Connection of queries about user.`,
	Aliases: []string{"u"},
	Example: `
  app user new
  app user paid
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
