package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/billylkc/app/calc"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// tProductsCmd represents the topProducts command
var tProductsCmd = &cobra.Command{
	Use:     "products",
	Short:   "Top selling products in monthly view.",
	Long:    `Top selling products in monthly view.`,
	Aliases: []string{"p"},
	Example: `
  app top products 10

  app t p 10
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var (
			nrecords  int
			headerM   map[string]bool
			headerR   table.Row
			tableRows []table.Row
			err       error
		)

		if len(args) == 0 {
			nrecords = 20
		} else {
			nrecords, err = strconv.Atoi(args[0])
			if err != nil {
				return err
			}
		}

		type Record struct {
			Date  string
			Sales int
		}

		topProducts, err := calc.GetTopProducts(nrecords)
		if err != nil {
			fmt.Println(err.Error())
		}
		headerM = make(map[string]bool)
		for _, v := range topProducts {
			for _, vv := range v {
				month := vv.Month
				if _, ok := headerM[month]; !ok {
					headerM[month] = true
				}
			}
		}
		// Sort the header slice
		var headerSorted []string
		for k, _ := range headerM {
			headerSorted = append(headerSorted, k)
		}
		sort.Strings(headerSorted)

		// Prepare header
		headerR = append(headerR, "Product")
		headerR = append(headerR, "GrandTotal")
		for _, k := range headerSorted {
			headerR = append(headerR, k)
		}

		for i := 1; i <= len(topProducts); i++ {
			var tr table.Row
			m := make(map[string]float64)
			rows := topProducts[i]

			username := rows[0].Field
			grandTotal := rows[0].GrandTotal
			tr = append(tr, username)
			tr = append(tr, humanize.CommafWithDigits(grandTotal, 1))
			for _, r := range rows {
				m[r.Month] = r.Total
			}
			for _, k := range headerSorted {
				if v, ok := m[k]; ok {
					vv := humanize.CommafWithDigits(v, 1)
					tr = append(tr, vv)
				} else {
					tr = append(tr, "")
				}
			}
			tableRows = append(tableRows, tr)
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(headerR)
		rowConfigAutoMerge := table.RowConfig{AutoMerge: false}
		t.AppendRows(tableRows, rowConfigAutoMerge)
		t.AppendSeparator()
		t.Style().Options.SeparateRows = true
		t.Render()

		return nil
	},
}

func init() {
	tCmd.AddCommand(tProductsCmd)
}
