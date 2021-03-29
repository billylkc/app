package calc

import (
	"fmt"
	"sort"
	"time"

	"github.com/billylkc/app/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/now"
	"github.com/pkg/errors"
)

// SalesRecord as daily sales for
type SalesRecord struct {
	Date  time.Time
	Count int
	Total float64
}

// GetDailySales returns the latest sales record for the past n days
func GetDailySales(d string, n int) ([]SalesRecord, error) {
	var records []SalesRecord

	// handle stupid date, add one day before query
	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return records, err
	}
	d = t.AddDate(0, 0, 1).Format("2006-01-02")

	db, err := database.GetConnection()
	if err != nil {
		return records, err
	}

	queryF := `
    SELECT
        DATE(created_date) as date,
        count(1) as count,
        SUM(total) as total
    FROM
        %s
    WHERE
        created_date <= '%s'
    GROUP BY
        DATE(created_date)
    ORDER BY
        created_date DESC
    LIMIT %d
    `
	// query := fmt.Sprintf(queryF, "`order`", `"2021-03-25"`, n)
	query := fmt.Sprintf(queryF, "`order`", d, n)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return records, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var rec SalesRecord
		err = results.Scan(&rec.Date, &rec.Count, &rec.Total)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		records = append(records, rec)

	}

	// Sort the stupid slice, desc
	sort.Slice(records, func(i, j int) bool {
		return records[j].Date.Before(records[i].Date)
	})

	return records, nil
}

// GetMonthsSales returns the latest sales record for the past n days
func GetMonthlySales(d string, n int) ([]SalesRecord, error) {
	var records []SalesRecord

	// Parse start end date
	t, err := now.Parse(d)
	if err != nil {
		return records, err
	}
	end := now.With(t).EndOfMonth().Format("2006-01-02")
	start := t.AddDate(0, -n+1, 0).Format("2006-01-02")

	db, err := database.GetConnection()
	if err != nil {
		return records, err
	}

	queryF := `
SELECT
	DATE,
	count(1) as count,
	SUM(total) as total
FROM (
SELECT
	%s,
	total
FROM
	%s
WHERE
	created_date >= '%s' and created_date <= '%s'
) AS o
GROUP BY
	DATE
ORDER BY
	DATE DESC
`
	query := fmt.Sprintf(queryF,
		"CAST(DATE_FORMAT(created_date,'%Y-%m-01') as DATE) as DATE",
		"`order`",
		start,
		end)

	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return records, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var rec SalesRecord
		err = results.Scan(&rec.Date, &rec.Count, &rec.Total)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		records = append(records, rec)
	}

	// Sort the stupid slice, desc
	sort.Slice(records, func(i, j int) bool {
		return records[j].Date.Before(records[i].Date)
	})

	return records, nil
}
