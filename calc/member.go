package calc

import (
	"fmt"
	"time"

	"github.com/billylkc/app/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/now"
	"github.com/pkg/errors"
)

// MemberRecord as sales record for different members
type MemberRecord struct {
	Date       time.Time
	ID         string
	Username   string
	Total      float64
	Average    float64
	GrandTotal float64
}

var memberLimit int

// GetDailyMember returns daily member spendings
func GetDailyMember(d string, n int) ([]MemberRecord, error) {
	var records []MemberRecord
	memberLimit = 50

	// handle stupid date, add one day before query
	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return records, err
	}
	start := t.AddDate(0, 0, -n+1).Format("2006-01-02") // start date
	end := t.AddDate(0, 0, 1).Format("2006-01-02")      // end date

	db, err := database.GetConnection()
	if err != nil {
		return records, err
	}

	queryF := `

SELECT
    DATE,
    op.CUSTOMER_ID,
    USERNAME,
    SUM(TOTAL) as TOTAL,
	AVERAGE,
	GrandTotal
FROM (

	SELECT
		CAST(op.created_date AS DATE) as DATE,
		op.CUSTOMER_ID,
		c.USERNAME,
		Total
	FROM order_product as op
		INNER JOIN customer as c
		  ON op.customer_id = c.id
	WHERE op.created_date >= '%s' and op.created_date <= '%s') as op

INNER JOIN

(
	SELECT
		CUSTOMER_ID,
		AVG(DayTotal) as AVERAGE,
		SUM(DayTotal) as GrandTotal
	FROM (

	SELECT
		DATE,
		CUSTOMER_ID,
		SUM(Total) as DayTotal
		FROM
			(SELECT
				CAST(created_date AS DATE) as DATE,
				CUSTOMER_ID,
				Total
			FROM order_product) as day
		GROUP BY DATE, CUSTOMER_ID
	) as o
	GROUP BY CUSTOMER_ID
) as tt
on op.CUSTOMER_ID = tt.CUSTOMER_ID
WHERE total >= %d
GROUP BY DATE, op.CUSTOMER_ID, USERNAME
ORDER BY DATE DESC, TOTAL DESC
    `
	query := fmt.Sprintf(queryF, start, end, memberLimit)

	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return records, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var rec MemberRecord
		err = results.Scan(&rec.Date, &rec.ID, &rec.Username, &rec.Total, &rec.Average, &rec.GrandTotal)

		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}
	return records, nil
}

// GetWeeklyMember returns weekly member spendings
func GetWeeklyMember(d string, n int) ([]MemberRecord, error) {
	var records []MemberRecord
	memberLimit = 200

	// handle stupid date, add one day before query
	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return records, err
	}
	end := now.With(t).Format("2006-01-02")
	start := t.AddDate(0, 0, -n*7).Format("2006-01-02")

	db, err := database.GetConnection()
	if err != nil {
		return records, err
	}

	queryF := `
SELECT
    DATE,
    oop.CUSTOMER_ID,
    USERNAME,
    TOTAL,
    AVERAGE,
    GrandTotal
FROM
        (SELECT
                DATE,
                CUSTOMER_ID,
                USERNAME,
                SUM(TOTAL) as TOTAL
        FROM
                (SELECT
						%s,
                        op.CUSTOMER_ID,
                        c.USERNAME,
                        Total
                FROM order_product as op
                        INNER JOIN customer as c
                          ON op.customer_id = c.id
                WHERE op.created_date >= '%s' and op.created_date <= '%s') as op
        GROUP BY
                DATE,
                CUSTOMER_ID,
                USERNAME) as oop

INNER JOIN
        (SELECT
        CUSTOMER_ID,
        AVG(MonthTotal) as AVERAGE,
        SUM(MonthTotal) as GrandTotal
    FROM (

                SELECT
                        DATE,
                        CUSTOMER_ID,
                        SUM(Total) as MonthTotal
                        FROM
                                (SELECT
                                        %s,
                                        CUSTOMER_ID,
                                        Total
                                FROM order_product) as day
                        GROUP BY DATE, CUSTOMER_ID
    ) as o
    GROUP BY CUSTOMER_ID
        ) as tt
on oop.CUSTOMER_ID = tt.CUSTOMER_ID
WHERE total >= %d
ORDER BY DATE DESC, TOTAL DESC
`
	query := fmt.Sprintf(queryF,
		"CAST(SUBDATE(op.created_date, WEEKDAY(op.created_date)) AS DATE) AS DATE",
		start,
		end,
		"CAST(SUBDATE(created_date, WEEKDAY(created_date)) AS DATE) AS DATE",
		memberLimit,
	)

	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return records, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var rec MemberRecord
		err = results.Scan(&rec.Date, &rec.ID, &rec.Username, &rec.Total, &rec.Average, &rec.GrandTotal)

		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}
	return records, nil
}

// GetMonthlyMember returns daily member spendings
func GetMonthlyMember(d string, n int) ([]MemberRecord, error) {
	var records []MemberRecord
	memberLimit = 1000

	// handle stupid date, add one day before query
	t, err := time.Parse("2006-01-02", d)
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
    oop.CUSTOMER_ID,
    USERNAME,
    TOTAL,
    AVERAGE,
    GrandTotal
FROM
	(SELECT
		DATE,
		CUSTOMER_ID,
		USERNAME,
		SUM(TOTAL) as TOTAL
	FROM
		(SELECT
            %s,
			op.CUSTOMER_ID,
			c.USERNAME,
			Total
		FROM order_product as op
			INNER JOIN customer as c
			  ON op.customer_id = c.id
		WHERE op.created_date >= '%s' and op.created_date <= '%s') as op
	GROUP BY
		DATE,
		CUSTOMER_ID,
		USERNAME) as oop

INNER JOIN
	(SELECT
        CUSTOMER_ID,
        AVG(MonthTotal) as AVERAGE,
        SUM(MonthTotal) as GrandTotal
    FROM (

		SELECT
			ADATE,
			CUSTOMER_ID,
			SUM(Total) as MonthTotal
			FROM
				(SELECT
					%s,
					CUSTOMER_ID,
					Total
				FROM order_product) as day
			GROUP BY ADATE, CUSTOMER_ID
    ) as o
    GROUP BY CUSTOMER_ID
	) as tt
on oop.CUSTOMER_ID = tt.CUSTOMER_ID
WHERE total >= %d
ORDER BY DATE DESC, TOTAL DESC
`
	query := fmt.Sprintf(queryF,
		"CAST(DATE_FORMAT(op.created_date,'%Y-%m-01') as DATE) as DATE",
		start,
		end,
		"CAST(DATE_FORMAT(order_product.created_date,'%Y-%m-01') as DATE) as ADATE",
		memberLimit,
	)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return records, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var rec MemberRecord
		err = results.Scan(&rec.Date, &rec.ID, &rec.Username, &rec.Total, &rec.Average, &rec.GrandTotal)

		if err != nil {
			panic(err.Error())
		}
		records = append(records, rec)
	}
	return records, nil
}
