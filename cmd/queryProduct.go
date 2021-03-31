package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/app/util"
	"github.com/spf13/cobra"
)

// productCmd represents the product command
var productCmd = &cobra.Command{
	Use:     "product",
	Short:   "Query the performance of individual product.",
	Long:    `Query the performance of individual product.`,
	Aliases: []string{"p"},
	Example: `  app query product 180125"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			id  int
			err error
		)
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		if len(args) >= 1 {
			pid := args[0] // product id
			id, err = strconv.Atoi(pid)
			if err != nil {
				return err
			}

		}

		// Print product details
		products, err := calc.GetProductDetails(id)
		if err != nil {
			return err
		}

		data := util.InterfaceSlice(products)
		headers := []string{"ID", "Name", "URL", "Price", "Listed_Price", "Discount", "Active"}
		ignores := []string{"Code"}
		fmt.Printf("\nProduct Details\n\n")
		err = util.PrintTable(data, headers, ignores, 1)
		if err != nil {
			return err
		}

		// Print sales related records
		records, err := calc.GetProductSalesRecords(id)
		if err != nil {
			return err
		}
		data = util.InterfaceSlice(records)
		headers = []string{"Date", "ProductID", "Total", "Quantity"}
		ignores = []string{"Category", "ProductName"}

		fmt.Printf("\nWeekly Sales\n\n")
		err = util.PrintTable(data, headers, ignores, 2)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	qCmd.AddCommand(productCmd)
}
