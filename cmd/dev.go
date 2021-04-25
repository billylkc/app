package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "[ ] For development purpose.",
	Long:  `[ ] For development purpose.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := myutil.HandleDateArgs(&date, &nrecords, 7, args...)
		if err != nil {
			return err
		}

		d, err := myutil.ParseDateInput(date, "d")
		if err != nil {
			return err
		}

		start, end, err := myutil.ParseDateRange(d, nrecords, "d")
		fmt.Println(start, end)

		return nil
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
