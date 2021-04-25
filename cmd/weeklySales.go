package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// wSalesCmd represents the weeklyales command
var wSalesCmd = &cobra.Command{
	Use:     "sales [-d date] [previous] [nrecords]",
	Short:   "[s] Weekly sales for the last n weeks.",
	Long:    `[s] Weekly sales for the last n weeks.`,
	Aliases: []string{"s"},
	Example: `
  app weekly sales -d "2021-03-24"
  app w s
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := myutil.HandleDateArgs(&date, &nrecords, 4, args...)
		if err != nil {
			return err
		}

		d, err := myutil.ParseDateInput(date, "w")
		if err != nil {
			return err
		}

		start, end, err := myutil.ParseDateRange(d, nrecords, "w")
		res, err := calc.GetWeeklySales(start, end)
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
	wCmd.AddCommand(wSalesCmd)
	wSalesCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
