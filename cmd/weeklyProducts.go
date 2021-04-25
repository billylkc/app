package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// wProductsCmd represents the weeklyProducts command
var wProductCmd = &cobra.Command{
	Use:     "products [-d date] [previous] [nrecords]",
	Short:   "[p] Weekly products.",
	Long:    `[p] Weekly products.`,
	Aliases: []string{"p"},
	Example: `
  app weekly products -d "2020-03-25"
  app w p 1 2 (Starting from one day ago, return two records)
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := myutil.HandleDateArgs(&date, &nrecords, 1, args...)
		if err != nil {
			return err
		}

		d, err := myutil.ParseDateInput(date, "w")
		if err != nil {
			return err
		}

		start, end, err := myutil.ParseDateRange(d, nrecords, "w")
		res, err := calc.GetWeeklyProduct(start, end)
		if err != nil {
			return err
		}

		headers := []string{"Date", "Cateogry", "ID", "Product Name", "Total", "Quantity"}
		ignores := []string{""}
		data := myutil.InterfaceSlice(res)
		err = myutil.PrintTable(data, headers, ignores, 5)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	wCmd.AddCommand(wProductCmd)
	wProductCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
