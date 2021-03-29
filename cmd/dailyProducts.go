package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	util "github.com/billylkc/app/util"
	"github.com/spf13/cobra"
)

// dailyProductsCmd represents the dailyProducts command
var dailyProductCmd = &cobra.Command{
	Use:     "products",
	Short:   "Daily products.",
	Long:    `Daily products.`,
	Aliases: []string{"p"},
	Example: `  app daily products -d "2020-03-25"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			date = args[0]
		}

		d, err := util.ParseDateInput(date)
		if err != nil {
			return err
		}

		res, err := calc.GetDailyProduct(d, 0)
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
	dailyCmd.AddCommand(dailyProductCmd)
	dailyProductCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
