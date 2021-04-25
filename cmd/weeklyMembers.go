package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// wMembersCmd represents the weekly members command
var wMembersCmd = &cobra.Command{
	Use:     "members",
	Short:   "[w] Weekly member spending.",
	Long:    `[w] Weekly member spending.`,
	Aliases: []string{"m"},
	Example: `
  app weekly members -d "2021-03-01"
  app w m 0 2
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := myutil.HandleDateArgs(&date, &nrecords, 1, args...)
		if err != nil {
			return err
		}

		d, err := myutil.ParseDateInput(date, "w")
		if err != nil {
			return err
		}

		start, end, err := myutil.ParseDateRange(d, nrecords, "w")
		res, err := calc.GetWeeklyMember(start, end)
		if err != nil {
			return err
		}

		headers := []string{"Date", "Username", "Monthly_Total", "Average", "Grand_Total"}
		ignores := []string{"ID"}
		data := myutil.InterfaceSlice(res)
		err = myutil.PrintTable(data, headers, ignores, 1)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	wCmd.AddCommand(wMembersCmd)
	wMembersCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
