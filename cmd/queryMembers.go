package cmd

import (
	"fmt"
	"strings"

	"github.com/billylkc/app/calc"
	"github.com/billylkc/app/util"
	"github.com/spf13/cobra"
)

// membersCmd represents the members command
var membersCmd = &cobra.Command{
	Use:     "members [name]",
	Short:   "[m] Query member purchase history.",
	Long:    `[m] Query member purchase history.`,
	Aliases: []string{"m"},
	Example: `
  app query members Jimbelle
  app q m Jimbelle
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var s string
		if len(args) > 0 {
			s = strings.Join(args, " ")
		} else {
			return fmt.Errorf("not enough argumetns\n")
		}
		memberSpendings, detailHistory, err := calc.GetPurchaseHistory(s)
		if err != nil {
			return err
		}

		// Overall spendings
		headers := []string{"Date", "Username", "Count", "Total"}
		ignores := []string{"ProductID", "ProductName"}
		data := util.InterfaceSlice(memberSpendings)
		fmt.Printf("\nMember Spendings\n\n")
		err = util.PrintTable(data, headers, ignores, 2)
		if err != nil {
			return err
		}

		// Detailed items purchased
		headers = []string{"Date", "Username", "ProductID", "ProductName", "Total"}
		ignores = []string{}
		data = util.InterfaceSlice(detailHistory)
		fmt.Printf("\nDetailed History\n\n")
		err = util.PrintTable(data, headers, ignores, 2)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	qCmd.AddCommand(membersCmd)
}
