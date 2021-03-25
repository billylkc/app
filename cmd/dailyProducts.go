package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// dailyProductsCmd represents the dailyProducts command
var dailyProductCmd = &cobra.Command{
	Use:     "products",
	Short:   "Daily products",
	Long:    `Daily products`,
	Aliases: []string{"p"},
	Example: `  app daily products -d "2020-03-25"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("product called")
		return nil
	},
}

func init() {
	today := time.Now().Format("2006-01-02")
	dailyCmd.AddCommand(dailyProductCmd)
	dailyProductCmd.Flags().StringVarP(&date, "date", "d", today, "Start date of query")
}
