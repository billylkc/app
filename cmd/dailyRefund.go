package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// dRefundCmd represents the refund command
var dRefundCmd = &cobra.Command{
	Use:     "refund",
	Short:   "[r] Daily refund for the past 7 days.",
	Long:    `[r] Daily refund for the past 7 days.`,
	Aliases: []string{"r"},
	Example: `
  app daily refund -d "2021-03-24"
  app d r 0 5
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
		res, err := calc.GetDailyRefund(start, end)
		if err != nil {
			return err
		}

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
	dCmd.AddCommand(dRefundCmd)
	dRefundCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
