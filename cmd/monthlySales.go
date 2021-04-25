package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// monthlySalesCmd represents the monthlySales command
var mSalesCmd = &cobra.Command{
	Use:     "sales [-d date] [previous] [nrecords]",
	Short:   "[s] Monthly sales details.",
	Long:    `[s] Monthly sales details.`,
	Aliases: []string{"s"},
	Example: `
  app monthly sales -d "2021-03-24"
  app m s 0 1
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := myutil.HandleDateArgs(&date, &nrecords, 1, args...)
		if err != nil {
			return err
		}

		d, err := myutil.ParseDateInput(date, "m")
		if err != nil {
			return err
		}

		start, end, err := myutil.ParseDateRange(d, nrecords, "m")
		res, err := calc.GetMonthlySales(start, end)
		if err != nil {
			return err
		}

		headers := []string{"Date", "Count", "Total"}
		ignores := []string{}
		data := myutil.InterfaceSlice(res)
		err = myutil.PrintTable(data, headers, ignores, 3)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	mCmd.AddCommand(mSalesCmd)
	mSalesCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
