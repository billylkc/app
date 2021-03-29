package cmd

import (
	"fmt"
	"time"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/app/util"
	"github.com/spf13/cobra"
)

// wSalesCmd represents the weeklyales command
var wSalesCmd = &cobra.Command{
	Use:     "sales",
	Short:   "Weekly sales for the last n weeks.",
	Long:    `Weekly sales for the last n weeks.`,
	Aliases: []string{"s"},
	Example: `  app daily sales -d "2021-03-24"`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := util.HandleDateArgs(&date, &nrecords, 4, args...)
		if err != nil {
			return err
		}

		d, err := util.ParseDateInput(date, "w")
		if err != nil {
			return err
		}

		res, err := calc.GetWeeklySales(d, nrecords)
		if err != nil {
			return err
		}
		fmt.Println(res)

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
