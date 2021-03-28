package app

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
)

// ParseDateInput parse the input for past n days, or actual day string in YYYY-MM-DD format
func ParseDateInput(s string) (string, error) {
	var res string

	// Check if input is in YYYY-MM-DD format
	if len(s) == 10 {
		// TODO: Check for pattern
		return s, nil
	}

	// Convert to date
	d, err := strconv.Atoi(s)
	d = d - 1
	if err != nil {
		return res, fmt.Errorf("Invalid input for date. Need a date in YYYY-MM-DD format or number for past n days")
	}
	res = time.Now().AddDate(0, 0, -d).Format("2006-01-02")
	fmt.Println(res)
	return res, nil
}

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
