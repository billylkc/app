package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// dSalesCmd represents the dailySales command
var dSalesCmd = &cobra.Command{
	Use:     "sales [-d date] [previous] [nrecords]",
	Short:   "[s] Daily sales for the past 7 days.",
	Long:    `[s] Daily sales for the past 7 days.`,
	Aliases: []string{"s"},
	Example: `
  app daily sales -d "2021-03-24"
  app d s 0 5
`,
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
		res, err := calc.GetDailySales(start, end)
		if err != nil {
			return err
		}

		headers := []string{"Date", "Count", "Total"}
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
	today := time.Now().Format("2006-01-02")
	dCmd.AddCommand(dSalesCmd)
	dSalesCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
