package cmd

import (
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/app/util"
	"github.com/spf13/cobra"
)

// dSalesCmd represents the dailySales command
var dSalesCmd = &cobra.Command{
	Use:     "sales",
	Short:   "Daily sales for the last 7 days.",
	Long:    `Daily sales for the last 7 days.`,
	Aliases: []string{"s"},
	Example: `  app daily sales -d "2021-03-24"`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 1 {
			date = args[0]
		}

		d, err := util.ParseDateInput(date, "d")
		if err != nil {
			return err
		}

		res, err := calc.GetDailySales(d, 7)
		if err != nil {
			return err
		}

		headers := []string{"Date", "Count", "Total"}
		ignores := []string{}
		data := util.InterfaceSlice(res)
		err = util.PrintTable(data, headers, ignores, 3)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	dCmd.AddCommand(dSalesCmd)
	dSalesCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
