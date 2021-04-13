package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/app/util"
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

		err := util.HandleDateArgs(&date, &nrecords, 4, args...)
		if err != nil {
			return err
		}

		// As weekly will take the incomplete week, need to subtract one week to balance it out
		if nrecords >= 1 {
			nrecords -= 1
		}

		d, err := util.ParseDateInput(date, "w")
		if err != nil {
			return err
		}

		res, err := calc.GetWeeklySales(d, nrecords)
		if err != nil {
			return err
		}

		headers := []string{"Date", "Count", "Total"}
		ignores := []string{}
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
	wCmd.AddCommand(wSalesCmd)
	wSalesCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
