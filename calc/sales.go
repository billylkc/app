package calc

import (
	"fmt"
	"sort"
	"time"

	"github.com/billylkc/app/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// SalesRecord as daily sales for
type SalesRecord struct {
	Date  time.Time
	Count int
	Total float64
}

// GetDailySales returns the latest sales record for the past n days
func GetDailySales(start, end string) ([]SalesRecord, error) {
	var records []SalesRecord

	// handle stupid date, add one day before query
	t, err := time.Parse("2006-01-02", end)
	if err != nil {
		return records, err
	}

	end = t.AddDate(0, 0, 1).Format("2006-01-02")
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
        created_date >= '%s' AND
        created_date <= '%s'
    GROUP BY
        DATE(created_date)
    ORDER BY
        created_date DESC
    `
	query := fmt.Sprintf(queryF, "`order`", start, end)
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

// GetWeeklySales returns the latest sales record for the past n weeks
func GetWeeklySales(start, end string) ([]SalesRecord, error) {
	var records []SalesRecord

	// handle stupid date, add one day before query
	t, err := time.Parse("2006-01-02", end)
	if err != nil {
		return records, err
	}

	end = t.AddDate(0, 0, 1).Format("2006-01-02")

	db, err := database.GetConnection()
	if err != nil {
		return records, err
	}

	queryF := `
SELECT
    DATE,
	COUNT(1),
	SUM(TOTAL) as TOTAL
FROM
    (SELECT
		CAST(SUBDATE(created_date, WEEKDAY(created_date)) AS DATE) AS DATE,
        TOTAL
    FROM %s as op
    WHERE
        created_date >= '%s' AND
        created_date <= '%s'
	) as oop
GROUP BY
	DATE
ORDER BY DATE DESC
	`

	query := fmt.Sprintf(queryF,
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

// GetMonthlySales returns the latest sales record for the past n months
func GetMonthlySales(start, end string) ([]SalesRecord, error) {
	var records []SalesRecord

	// handle stupid date, add one day before query
	t, err := time.Parse("2006-01-02", end)
	if err != nil {
		return records, err
	}
	end = t.AddDate(0, 0, 1).Format("2006-01-02")

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
