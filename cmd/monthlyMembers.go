package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/app/util"
	"github.com/spf13/cobra"
)

// dMembersCmd represents the daily members command
var mMembersCmd = &cobra.Command{
	Use:     "members",
	Short:   "Monthly member spending.",
	Long:    `Monthly member spending.`,
	Aliases: []string{"m"},
	Example: `  app monthly members -d "2021-03-01"`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := util.HandleDateArgs(&date, &nrecords, 1, args...)
		if err != nil {
			return err
		}

		d, err := util.ParseDateInput(date, "m")
		if err != nil {
			return err
		}

		res, err := calc.GetMonthlyMember(d, nrecords)
		if err != nil {
			return err
		}
		headers := []string{"Date", "Username", "Monthly_Total", "Average", "Grand_Total"}
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
	mCmd.AddCommand(mMembersCmd)
	mMembersCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
