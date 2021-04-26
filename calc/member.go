package calc

import (
	"fmt"
	"strings"
	"time"

	"github.com/billylkc/app/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/sahilm/fuzzy"
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

type PurchaseHistory struct {
	Date        time.Time
	Username    string
	ProductID   int
	ProductName string
	Count       int
	Total       float64
}

type MonthlySales struct {
	Field      string // username, product_id
	Name       string //  productName
	GrandTotal float64
	Month      string
	Total      float64
}

type members []MemberRecord

var memberLimit int // Filtering criteria for listing members

func (m members) String(i int) string {
	return m[i].Username
}

func (m members) Len() int {
	return len(m)
}

// GetDailyMember returns daily member spendings
func GetDailyMember(start, end string) ([]MemberRecord, error) {
	var records []MemberRecord
	memberLimit = 50

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
func GetWeeklyMember(start, end string) ([]MemberRecord, error) {
	var records []MemberRecord
	memberLimit = 200

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

// GetMonthlyMember returns monthly member spendings
func GetMonthlyMember(start, end string) ([]MemberRecord, error) {
	var records []MemberRecord
	memberLimit = 1000

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

// GetPurchaseHistory gets the monthly purchase history of the member queried
// member name is matched with fuzzy search (loosely)
func GetPurchaseHistory(s string) ([]PurchaseHistory, []PurchaseHistory, error) {
	var (
		md   []PurchaseHistory // Member details - member monthly total, etc
		ph   []PurchaseHistory // Purchase history - items per month
		full members           // Use by fuzzy search
		ids  []string          // list of members id result from fuzzy search
	)

	db, err := database.GetConnection()
	if err != nil {
		return md, ph, err
	}

	// Get list of user id first using fuzzy search
	queryF := `
    SELECT
        ID,
        USERNAME
    FROM customer
    `
	query := fmt.Sprintf(queryF)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return md, ph, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var rec MemberRecord
		err = results.Scan(&rec.ID, &rec.Username)

		if err != nil {
			panic(err.Error())
		}
		full = append(full, rec)
	}
	res := fuzzy.FindFrom(s, full)
	for _, r := range res {
		ids = append(ids, fmt.Sprintf("\"%s\"", full[r.Index].ID))
	}
	idList := strings.Join(ids, ", ")

	// Member total
	queryF = `
    SELECT
        date,
        c.username,
        count(1) as count,
        sum(total) as total
     FROM (
        SELECT
            %s,
            customer_id,
            total
        FROM
            order_product
        WHERE
            customer_id in (%s)
     ) as op
         INNER JOIN
             customer as c
                 on c.id = op.customer_id
     GROUP BY date, customer_id
     order by customer_id, date desc, total desc
`
	query = fmt.Sprintf(queryF,
		"CAST(DATE_FORMAT(created_date,'%Y-%m-01') as DATE) as date",
		idList,
	)

	results, err = db.Query(query)
	defer results.Close()
	if err != nil {
		return md, ph, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var rec PurchaseHistory
		err = results.Scan(&rec.Date, &rec.Username, &rec.Count, &rec.Total)
		if err != nil {
			panic(err.Error())
		}
		md = append(md, rec)
	}

	// Detailed purchase history
	queryF = `
    SELECT
        date,
        c.username,
        op.product_id,
        p.name as product_name,
        count(1) as count,
        sum(total) as total
     FROM (
        SELECT
            %s,
            customer_id,
            product_id,
            total
        FROM
            order_product
        WHERE
            customer_id in (%s)
     ) as op
         INNER JOIN
             product as p
                 on p.product_id = op.product_id
         INNER JOIN
             customer as c
                 on c.id = op.customer_id
     GROUP BY date, customer_id, product_id
     order by customer_id, date desc, total desc
    `
	query = fmt.Sprintf(queryF,
		"CAST(DATE_FORMAT(created_date,'%Y-%m-01') as DATE) as date",
		idList,
	)

	results, err = db.Query(query)
	defer results.Close()

	if err != nil {
		return md, ph, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var rec PurchaseHistory
		err = results.Scan(&rec.Date, &rec.Username, &rec.ProductID, &rec.ProductName, &rec.Count, &rec.Total)

		if err != nil {
			panic(err.Error())
		}
		ph = append(ph, rec)
	}

	return md, ph, nil
}

func GetTopMembers(nrecords int) (map[int][]MonthlySales, error) {
	m := make(map[int][]MonthlySales)

	db, err := database.GetConnection()
	if err != nil {
		return m, err
	}

	queryF := `
SELECT
	RANK,
	c.username,
	GrandTotal,
	YearMonth,
	MonthTotal
FROM
(
	SELECT
		%s,
		CUSTOMER_ID,
		SUM(Total) as MonthTotal
	FROM
			(SELECT
					%s,
					CUSTOMER_ID,
					Total
			FROM order_product) as day
	GROUP BY DATE, CUSTOMER_ID
) as op

INNER JOIN

(SELECT @rn:=@rn+1 AS rank, customer_id, GrandTotal
FROM (
  SELECT
		customer_id,
		sum(total) as GrandTotal
	FROM
		order_product
	WHERE
		customer_id is not null
	GROUP BY
		customer_id
	ORDER BY
		GrandTotal desc
) t1, (SELECT @rn:=0) t2) as tt

	on op.customer_id = tt.customer_id

INNER JOIN
	customer as c
		on c.id = op.customer_id

WHERE tt.rank <= %d
order by rank, YearMonth desc, MonthTotal desc
    `
	query := fmt.Sprintf(queryF,
		"DATE_FORMAT(DATE,'%Y-%m') as YearMonth",
		"CAST(DATE_FORMAT(order_product.created_date,'%Y-%m-01') as DATE) as DATE",
		nrecords,
	)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return m, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var (
			rank int
			rec  MonthlySales
		)
		err = results.Scan(&rank, &rec.Field, &rec.GrandTotal, &rec.Month, &rec.Total)
		if v, ok := m[rank]; ok {
			m[rank] = append(v, rec)
		} else {
			m[rank] = []MonthlySales{rec}
		}
	}
	return m, nil
}
