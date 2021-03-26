package cmd

import (
	"os"
	"time"

	"github.com/billylkc/app/calc"
	"github.com/jedib0t/go-pretty/v6/table"
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
		res, err := calc.GetDailyProduct(date, 0)
		if err != nil {
			return err
		}

		// Display table
		rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Date", "Cateogry", "ID", "Product Name", "Quantity", "Total"})
		for _, r := range res {
			date := r.Date.Format("2006-01-02")
			t.AppendRow(table.Row{date, r.Category, r.ProductID, r.ProductName, r.Quantity, r.Total}, rowConfigAutoMerge)
		}
		t.SetColumnConfigs([]table.ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
		})
		t.AppendSeparator()
		t.Style().Options.SeparateRows = true
		t.Render()

		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	dailyCmd.AddCommand(dailyProductCmd)
	dailyProductCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
