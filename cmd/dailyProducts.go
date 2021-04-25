package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/app/util"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// dProductsCmd represents the dailyProducts command
var dProductCmd = &cobra.Command{
	Use:     "products [-d date] [previous] [nrecords]",
	Short:   "[p] Daily top selling products.",
	Long:    `[p] Daily top selling products.`,
	Aliases: []string{"p"},
	Example: `
  app daily products -d "2020-03-25"
  app d p 1 2 (Starting from one day ago, return two records)
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := myutil.HandleDateArgs(&date, &nrecords, 1, args...)
		if err != nil {
			return err
		}

		d, err := myutil.ParseDateInput(date, "d")
		if err != nil {
			return err
		}

		start, end, err := myutil.ParseDateRange(d, nrecords, "d")
		res, err := calc.GetDailyProduct(start, end)
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
	dCmd.AddCommand(dProductCmd)
	dProductCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
