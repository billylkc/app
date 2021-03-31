package util

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jinzhu/now"
)

// HandleDateArgs handles non flag input arguments
// mostly handle nrecords only, where n could be days, weeks, or months
func HandleDateArgs(date *string, nrecords *int, defaultN int, args ...string) error {
	var err error
	if len(args) == 1 {
		*date = args[0]
		*nrecords = defaultN
	} else if len(args) == 2 {
		*date = args[0]
		*nrecords, err = strconv.Atoi(args[1])
		if err != nil {
			return err
		}
	} else {
		*nrecords = defaultN
	}
	return nil
}

// ParseDateInput parse the input for past n days, or actual day string in YYYY-MM-DD format
// result depends on freq, daily, monthly -> 2021-03-01, weekly -> start from monday
func ParseDateInput(s, freq string) (string, error) {
	var (
		t     time.Time
		dateF string
		err   error
	)

	// Set config
	location, _ := time.LoadLocation("Asia/Shanghai")
	tconfig := &now.Config{
		WeekStartDay: time.Monday,
		TimeLocation: location,
		TimeFormats:  []string{"2006-01-02"},
	}

	// Check if input is in YYYY-MM-DD format
	if len(s) == 10 {
		t, err = now.Parse(s)
		if err != nil {
			return dateF, fmt.Errorf("Invalid input for date. Need a date in YYYY-MM-DD format or number for past n days/weeks/months.")
		}
	} else {
		// Convert to date
		n, err := strconv.Atoi(s)
		if err != nil {
			return dateF, fmt.Errorf("Invalid input for date. Need a date in YYYY-MM-DD format or number for past n days/weeks/months")
		}
		switch freq {
		case "d": // Daily
			t = time.Now().AddDate(0, 0, -n)

		case "w": // Daily
			t = time.Now().AddDate(0, 0, -n*7)

		case "m": //Monthly
			t = time.Now().AddDate(0, -n, 0)

		default:
			t = time.Now()
		}
	}

	// Handle frequency
	switch freq {
	case "d": // Daily
		dateF = t.Format("2006-01-02")

	case "w": // Weekly
		dateF = tconfig.With(t).BeginningOfWeek().Format("2006-01-02")

	case "m": // Monthly
		dateF = tconfig.With(t).BeginningOfMonth().Format("2006-01-02")

	default:
		dateF = time.Now().Format("2006-01-02")
	}

	return dateF, nil
}

// InterfaceSlice converts a list of struct to list of interface
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		// panic("InterfaceSlice() given a non-slice type")
		ret := make([]interface{}, 1)
		ret = append(ret, slice)
		return ret
	}

	ret := make([]interface{}, s.Len())
	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

// PrintTable prints the table with interface
func PrintTable(data []interface{}, headers, ignore []string, colMerge int) error {

	// Build ignore column map from the input list
	// use to filter fields like ID from the result interface
	m := make(map[string]bool)
	for _, key := range ignore {
		if _, ok := m[key]; !ok {
			m[key] = true
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Handle headers
	var headersR table.Row
	for _, h := range headers {
		headersR = append(headersR, h)
	}
	t.AppendHeader(headersR)

	// Construct rows
	var rows []table.Row
	for _, x := range data {
		var r table.Row
		values := reflect.ValueOf(x)
		types := reflect.TypeOf(x)
		for i := 0; i < values.NumField(); i++ {
			field := values.Field(i).Interface()
			name := types.Field(i).Name

			// if name in the ignore map, skip
			if _, ok := m[name]; ok {
				continue
			}
			switch t := field.(type) {
			case time.Time: // for time, format to YYYY-MM-DD
				v := t.Format("2006-01-02")
				r = append(r, v)
			case float64:
				v := humanize.CommafWithDigits(t, 1)
				r = append(r, v)
			case int64:
				v := ""
				if strings.Contains(name, "ID") { // Handle ID column, no comma
					v = fmt.Sprintf("%d", t)
				} else {
					v = humanize.Comma(t)
				}
				r = append(r, v)
			case int:
				v := ""
				if strings.Contains(name, "ID") { // Handle ID column, no comma
					v = fmt.Sprintf("%d", t)
				} else {
					v = humanize.Comma(int64(t))
				}
				r = append(r, v)
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
