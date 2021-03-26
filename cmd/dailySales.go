package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/billylkc/app/calc"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// dailySalesCmd represents the dailySales command
var dailySalesCmd = &cobra.Command{
	Use:     "sales",
	Short:   "Daily sales for the last 7 days.",
	Long:    `Daily sales for the last 7 days.`,
	Aliases: []string{"s"},
	Example: `  app daily sales -d "2021-03-24"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := calc.GetDailySales(date, 7)
		if err != nil {
			fmt.Println(err)
		}

		// Display table
		rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Date", "Count", "Total"})
		for _, r := range res {
			date := r.Date.Format("2006-01-02")
			t.AppendRow(table.Row{date, r.Count, r.Total}, rowConfigAutoMerge)
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
	dailyCmd.AddCommand(dailySalesCmd)
	dailySalesCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
