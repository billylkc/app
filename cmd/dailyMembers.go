package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/app/util"
	"github.com/spf13/cobra"
)

// dMembersCmd represents the daily members command
var dMembersCmd = &cobra.Command{
	Use:     "members [-d date] [previous] [nrecords]",
	Short:   "[m] Daily member spending.",
	Long:    `[m] Daily member spending.`,
	Aliases: []string{"m"},
	Example: `
  app daily members -d "2021-03-25"
  app d m 0 1
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := util.HandleDateArgs(&date, &nrecords, 1, args...)
		if err != nil {
			return err
		}

		d, err := util.ParseDateInput(date, "d")
		if err != nil {
			return err
		}

		res, err := calc.GetDailyMember(d, nrecords)
		if err != nil {
			return err
		}
		headers := []string{"Date", "Username", "Day_Total", "Average", "Grand_Total"}
		ignores := []string{"ID"}
		data := util.InterfaceSlice(res)
		err = util.PrintTable(data, headers, ignores, 1)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	dCmd.AddCommand(dMembersCmd)
	dMembersCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
