package cmd

import (
	"github.com/billylkc/app/calc"
	"github.com/billylkc/myutil"
	"github.com/spf13/cobra"
)

// userPaidCmd represents the userPaid command
var userPaidCmd = &cobra.Command{
	Use:     "paid",
	Short:   "[p] Show no of paying users by month.",
	Long:    `[p] Show no of paying users by month.`,
	Aliases: []string{"p"},
	Example: `
  app user paid
  app u p 1 1
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := myutil.HandleDateArgs(&date, &nrecords, 4, args...)
		if err != nil {
			return err
		}
		d, err := myutil.ParseDateInput(date, "m")
		if err != nil {
			return err
		}

		start, end, err := myutil.ParseDateRange(d, nrecords, "m")
		if err != nil {
			return err
		}

		res, err := calc.GetPaidUserCount(start, end)
		if err != nil {
			return err
		}
		headers := []string{"Date", "Unique Members"}
		ignores := []string{"Country"}
		data := myutil.InterfaceSlice(res)
		err = myutil.PrintTable(data, headers, ignores, 1)
		if err != nil {
			return err
		}
		return nil

	},
}

func init() {
	userCmd.AddCommand(userPaidCmd)
}
