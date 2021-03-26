package calc

import (
	"fmt"
	"time"

	"github.com/billylkc/app/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// ProductRecord as sales record for differernt products
type ProductRecord struct {
	Date        time.Time
	Category    string
	ProductID   int
	ProductName string
	Quantity    int
	Total       float64
}

// GetDailyProduct returns the latest product sales record for the past n days
func GetDailyProduct(d string, n int) ([]ProductRecord, error) {
	var records []ProductRecord

	// handle stupid date, add one day before query
	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return records, err
	}
	start := t.AddDate(0, 0, -n).Format("2006-01-02") // start date
	end := t.AddDate(0, 0, 1).Format("2006-01-02")    // end date

	db, err := database.GetConnection()
	if err != nil {
		return records, err
	}

	queryF := `
    SELECT
        op.OrderDate as Date,
        COALESCE(pc.category_name, ""),
        op.product_id,
        p.product_name,
        SUM(op.quantity) quantiy,
        SUM(op.total) total
    FROM (
        SELECT
            created_date as OrderDate,
            product_id,
            quantity,
            total
        FROM
            order_product
        ORDER BY
            ORDER_PRODUCT_ID DESC
    ) as op

    LEFT OUTER JOIN

    (
        SELECT
               product_id,
               name as product_name
        FROM
               product
    ) as p

    ON op.product_id = p.product_id

    LEFT OUTER JOIN

    (SELECT
        pc.product_id,
        pc.category_id,
        c.name as category_name
    FROM product_to_category as pc
        INNER JOIN category as c
          on pc.category_id = c.category_id
    WHERE c.language_id = 1) as pc

    ON pc.product_id = op.product_id
    WHERE op.OrderDate >= '%s' and op.OrderDate <= '%s'
    GROUP BY op.OrderDate, pc.category_name, op.product_id, p.product_name
    ORDER BY OrderDate desc, pc.category_name, op.product_id
    `
	query := fmt.Sprintf(queryF, start, end)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return records, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var rec ProductRecord
		err = results.Scan(&rec.Date, &rec.Category, &rec.ProductID, &rec.ProductName, &rec.Quantity, &rec.Total)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		records = append(records, rec)
	}
	return records, nil
}
