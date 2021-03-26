package cmd

import (
	"os"
	"time"

	"github.com/billylkc/app/calc"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// dailyMembersCmd represents the dailyMembers command
var dailyMembersCmd = &cobra.Command{
	Use:     "members",
	Short:   "Daily member spending.",
	Long:    `Daily member spending.`,
	Aliases: []string{"m"},
	Example: `  app daily members -d "2021-03-25"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := calc.GetDailyMember(date, 0)
		if err != nil {
			return err
		}

		// Display table
		rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Date", "Username", "Total", "Average", "GrandTotal"})
		for _, r := range res {
			date := r.Date.Format("2006-01-02")
			t.AppendRow(table.Row{date, r.Username, humanize.Comma(int64(r.Total)), humanize.Comma(int64(r.Average)), humanize.Comma(int64(r.GrandTotal))}, rowConfigAutoMerge)
		}
		t.SetColumnConfigs([]table.ColumnConfig{
			{Number: 1, AutoMerge: true},
		})
		t.AppendSeparator()
		t.Style().Options.SeparateRows = true
		t.Render()

		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	dailyCmd.AddCommand(dailyMembersCmd)
	dailyMembersCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
