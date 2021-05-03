package cmd

import (
	"fmt"
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// mRefundCmd represents the refund command
var mRefundCmd = &cobra.Command{
	Use:     "refund",
	Short:   "[r] Monthly refund for the past 7 days.",
	Long:    `[r] Monthly refund for the past 7 days.`,
	Aliases: []string{"r"},
	Example: `
  app monthly refund -d "2021-03-24"
  app m r 0 5
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
		res, err := calc.GetMonthlyRefund(start, end)
		if err != nil {
			return err
		}
		fmt.Println(start)
		fmt.Println(end)

		headers := []string{"Date", "Count", "Total"}
		ignores := []string{}
		data := myutil.InterfaceSlice(res)
		err = myutil.PrintTable(data, headers, ignores, 0)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	mCmd.AddCommand(mRefundCmd)
	mRefundCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
