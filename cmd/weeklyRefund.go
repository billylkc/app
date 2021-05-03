package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// wRefundCmd represents the refund command
var wRefundCmd = &cobra.Command{
	Use:     "refund",
	Short:   "[r] Weekly refund for the last n weeks.",
	Long:    `[r] Weekly refund for the last n weeks.`,
	Aliases: []string{"r"},
	Example: `
  app weekly refund -d "2021-03-24"
  app w r 0 5
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
		res, err := calc.GetWeeklyRefund(start, end)
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
	wCmd.AddCommand(wRefundCmd)
	wRefundCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
