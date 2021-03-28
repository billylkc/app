package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	util "github.com/billylkc/app/util"
	"github.com/spf13/cobra"
)

// dailySalesCmd represents the dailySales command
var dailySalesCmd = &cobra.Command{
	Use:     "sales",
	Short:   "Daily sales for the last 7 days.",
	Long:    `Daily sales for the last 7 days.`,
	Aliases: []string{"s"},
	Example: `  app daily sales -d "2021-03-24"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		d, err := util.ParseDateInput(date)
		if err != nil {
			return err
		}

		res, err := calc.GetDailySales(d, 7)
		if err != nil {
			return err
		}

		headers := []string{"Date", "Count", "Total"}
		data := util.InterfaceSlice(res)
		err = util.PrintTable(data, headers, 3)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	dailyCmd.AddCommand(dailySalesCmd)
	dailySalesCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
