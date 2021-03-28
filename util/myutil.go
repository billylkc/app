package app

import (
	"os"
	"reflect"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
)

// InterfaceSlice converts a list of struct to list of interface
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

// PrintTable prints the table with interface
func PrintTable(data []interface{}, headers []string, colMerge int) error {
	var rows []table.Row
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Handle headers
	var headersR table.Row
	for _, h := range headers {
		headersR = append(headersR, h)
	}
	t.AppendHeader(headersR)

	// Construct rows
	for _, x := range data {
		var r table.Row
		v := reflect.ValueOf(x)
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i).Interface()
			switch t := field.(type) {
			case time.Time: // for time, format to YYYY-MM-DD
				value := t.Format("2006-01-02")
				r = append(r, value)
			case float64:
				value := humanize.CommafWithDigits(t, 1)
				r = append(r, value)
			case int64:
				value := humanize.Comma(t)
				r = append(r, value)
			case int:
				value := humanize.Comma(int64(t))
				r = append(r, value)
			default:
				r = append(r, field)
			}
		}
		rows = append(rows, r)
	}

	// Generate merged column config
	var cc []table.ColumnConfig
	if colMerge >= 1 {
		for i := 1; i <= colMerge; i++ {
			c := table.ColumnConfig{
				Number:    i,
				AutoMerge: true,
			}
			cc = append(cc, c)
		}
	}

	// Append rows
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t.AppendRows(rows, rowConfigAutoMerge)
	t.SetColumnConfigs(cc)
	t.AppendSeparator()
	t.Style().Options.SeparateRows = true
	t.Render()
	return nil
}
