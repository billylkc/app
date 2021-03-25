package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/billylkc/app/app"
	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "For development purpose.",
	Long:  `For development purpose.`,
	Run: func(cmd *cobra.Command, args []string) {
		res := app.Dev()
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
