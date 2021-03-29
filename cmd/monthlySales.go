package cmd

import (
	"fmt"
	"time"

	"github.com/billylkc/app/util"
	"github.com/spf13/cobra"
)

// monthlySalesCmd represents the monthlySales command
var mSalesCmd = &cobra.Command{
	Use:     "sales",
	Short:   "Monthly sales details.",
	Long:    `Monthly sales details.`,
	Aliases: []string{"s"},
	Example: `  app monthly sales -d "2021-03-24"`,
	RunE: func(cmd *cobra.Command, args []string) error {

		err := util.HandleDateArgs(&date, &nrecords, 1, args...)
		if err != nil {
			return err
		}

		d, err := util.ParseDateInput(date, "m")
		if err != nil {
			return err
		}
		fmt.Println(d)
		// res, err := calc.GetMonthlySales(d, 1)
		// if err != nil {
		// 	return err
		// }

		// headers := []string{"Date", "Count", "Total"}
		// ignores := []string{}
		// data := util.InterfaceSlice(res)
		// err = util.PrintTable(data, headers, ignores, 3)
		// if err != nil {
		// 	return err
		// }
		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	mCmd.AddCommand(mSalesCmd)
	mSalesCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
