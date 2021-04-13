package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/app/util"
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

		err := util.HandleDateArgs(&date, &nrecords, 1, args...)
		if err != nil {
			return err
		}

		if nrecords >= 1 {
			nrecords -= 1
		}

		d, err := util.ParseDateInput(date, "w")
		if err != nil {
			return err
		}

		res, err := calc.GetWeeklyProduct(d, nrecords)
		if err != nil {
			return err
		}

		headers := []string{"Date", "Cateogry", "ID", "Product Name", "Total", "Quantity"}
		ignores := []string{""}
		data := util.InterfaceSlice(res)
		err = util.PrintTable(data, headers, ignores, 5)
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
