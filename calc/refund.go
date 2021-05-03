package calc

import (
	"fmt"
	"sort"
	"time"

	"github.com/billylkc/app/database"
	"github.com/billylkc/app/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/now"
	"github.com/pkg/errors"
)

// RefundRecord as a struct for refund record
type RefundRecord struct {
	Date  time.Time
	Count int
	Total float64
}

// GetDailyRefund returns the latest refund record for the past n days
func GetDailyRefund(start, end string) ([]RefundRecord, error) {
	var records []RefundRecord

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
	DATE(od.created_time) as date,
	count(1) as count,
	SUM(op.refund_amount) as total

FROM order_product op
	INNER JOIN %s o
		ON (o.order_id = op.order_id)
	INNER JOIN order_delivery od ON (od.delivery_sn = o.delivery_sn)

WHERE
	(od.created_time >= '%s') AND
	(od.created_time <= '%s') AND
	(op.refund_amount > 0) AND (od.is_active = 1)
GROUP BY
	DATE(od.created_time)
ORDER BY
	od.created_time DESC

    `
	query := fmt.Sprintf(queryF,
		"`order`",
		start,
		end,
	)

	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return records, errors.Wrap(err, "cant execute query")
	}

	var md map[string]bool // List of days with records
	md = make(map[string]bool)
	for results.Next() {
		var rec RefundRecord
		err = results.Scan(&rec.Date, &rec.Count, &rec.Total)
		if err != nil {
			return records, err
		}
		records = append(records, rec)
		md[rec.Date.Format("2006-01-02")] = true
	}

	// Fill dates with empty records
	ss, err := util.GenerateDate(start, end, "d")
	if err != nil {
		return records, err
	}
	for _, s := range ss {
		if _, ok := md[s]; !ok {
			tt := now.MustParse(s)
			r := RefundRecord{
				Date: tt,
			}
			records = append(records, r)
		}

	}

	// Sort the stupid slice, desc
	sort.Slice(records, func(i, j int) bool {
		return records[j].Date.Before(records[i].Date)
	})

	return records, nil
}

func GetWeeklyRefund(start, end string) ([]RefundRecord, error) {
	var records []RefundRecord
	return records, nil
}

func GetMonthlyRefund(start, end string) ([]RefundRecord, error) {
	var records []RefundRecord
	return records, nil
}
