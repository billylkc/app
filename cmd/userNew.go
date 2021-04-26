package cmd

import (
	"github.com/billylkc/app/calc"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// userNewCmd represents the userNew command
var userNewCmd = &cobra.Command{
	Use:     "new",
	Short:   "Show no of new users by week.",
	Long:    `Show no of new users by week.`,
	Aliases: []string{"n"},
	Example: `
  app user new
  app user country
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := myutil.HandleDateArgs(&date, &nrecords, 4, args...)
		if err != nil {
			return err
		}
		d, err := myutil.ParseDateInput(date, "m")
		if err != nil {
			return err
		}

		start, end, err := myutil.ParseDateRange(d, nrecords, "m")
		if err != nil {
			return err
		}

		res, err := calc.GetNewUserCount(start, end)
		if err != nil {
			return err
		}
		headers := []string{"Date", "Country", "Count"}
		ignores := []string{}
		data := myutil.InterfaceSlice(res)
		err = myutil.PrintTable(data, headers, ignores, 1)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	userCmd.AddCommand(userNewCmd)
}
