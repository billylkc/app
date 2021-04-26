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

// GetNewUserCount gets the total of new user count by country per each week
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
