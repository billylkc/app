package calc

import (
	"fmt"
	"sort"
	"time"

	"github.com/billylkc/app/database"
	"github.com/pkg/errors"
)

// MemberCont
type MemberCount struct {
	Date    time.Time
	Country string
	Count   int
}

// GetPaidUserCount gets paid user per month
func GetPaidUserCount(start, end string) ([]MemberCount, error) {
	var records []MemberCount

	// handle stupid date, add one day before query
	t, err := time.Parse("2006-01-02", end)
	if err != nil {
		return records, err
	}
	end = t.AddDate(0, 0, 1).Format("2006-01-02") // end date

	db, err := database.GetConnection()
	if err != nil {
		return records, err
	}

	queryF := `
SELECT
	DATE,
	COUNT(1) as Count
FROM
(
	SELECT
		DATE,
		CUSTOMER_ID,
		count(1) as count
	FROM (
	SELECT
		%s,
		customer_id
	FROM
		%s
	WHERE
		created_date >= '%s' and created_date <= '%s'

	) AS o
	GROUP BY
		DATE, CUSTOMER_ID
	ORDER BY
		DATE DESC
) as oo
GROUP BY DATE
ORDER BY DATE DESC
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
		var rec MemberCount
		err = results.Scan(&rec.Date, &rec.Count)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}

	// Sort the stupid slice, desc
	sort.Slice(records, func(i, j int) bool {
		return records[j].Date.Before(records[i].Date)
	})
	return records, nil
}

// GetNewUserCount gets the total of new user count by country per month
func GetNewUserCount(start, end string) ([]MemberCount, error) {
	var records []MemberCount

	db, err := database.GetConnection()
	if err != nil {
		return records, err
	}

	queryF := `
SELECT
    c.DATE,
	l.SHORT_NAME AS COUNTRY,
	COUNT(1) AS COUNT
FROM
    (SELECT
        %s,
        USERNAME,
		LANGUAGE_ID
    FROM CUSTOMER
    ) as c
		INNER JOIN
	LANGUAGE as l
		ON c. LANGUAGE_ID = l.LANGUAGE_ID
WHERE DATE >= '%s' AND DATE <= '%s'
GROUP BY
    DATE, l.SHORT_NAME
ORDER BY DATE DESC, l.SHORT_NAME
`
	query := fmt.Sprintf(queryF,
		"CAST(DATE_FORMAT(created_date,'%Y-%m-01') as DATE) as DATE",
		start,
		end)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return records, errors.Wrap(err, "cant execute query")
	}
	for results.Next() {
		var rec MemberCount
		err = results.Scan(&rec.Date, &rec.Country, &rec.Count)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}

	// Sort the stupid slice, desc
	sort.Slice(records, func(i, j int) bool {
		return records[j].Date.Before(records[i].Date)
	})
	return records, nil
}
