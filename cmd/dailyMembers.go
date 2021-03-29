package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	util "github.com/billylkc/app/util"
	"github.com/spf13/cobra"
)

// dailyMembersCmd represents the dailyMembers command
var dailyMembersCmd = &cobra.Command{
	Use:     "members",
	Short:   "Daily member spending.",
	Long:    `Daily member spending.`,
	Aliases: []string{"m"},
	Example: `  app daily members -d "2021-03-25"`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 1 {
			date = args[0]
		}

		d, err := util.ParseDateInput(date)
		if err != nil {
			return err
		}

		res, err := calc.GetDailyMember(d, 0)
		if err != nil {
			return err
		}
		headers := []string{"Date", "Username", "DayTotal", "Average", "GrandTotal"}
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
	dailyCmd.AddCommand(dailyMembersCmd)
	dailyMembersCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
