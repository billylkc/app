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
	Total       float64
	Quantity    int
}

// Product as the product details of a particular item
type Product struct {
	ID          int     // product_id
	Code        string  // code
	Name        string  // Name
	URL         string  // url
	Price       float64 // price
	ListedPrice float64 // tb_list_price
	Discount    float64 // price
	Active      bool
}

var totalLimit int // Limit for printing popular Product

// GetDailyProduct returns the latest product sales record for the past n days
func GetDailyProduct(start, end string) ([]ProductRecord, error) {
	var records []ProductRecord
	totalLimit = 100
	fmt.Printf("\nShowing daily top products with sales larger than - %d \n\n", totalLimit)

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
*
FROM
(
    SELECT
        op.OrderDate as Date,
        COALESCE(pc.category_name, ""),
        op.product_id,
        p.product_name,
        SUM(op.quantity) quantiy,
        SUM(op.total) lumpsum
    FROM (
        SELECT
            CAST(created_date AS DATE) as OrderDate,
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
    ORDER BY OrderDate desc, total desc, pc.category_name, op.product_id
) as final
WHERE lumpsum >= %d
ORDER BY Date desc, lumpsum desc
    `
	query := fmt.Sprintf(queryF, start, end, totalLimit)
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

// GetWeeklyProduct returns the latest product sales record for the past n weeks
func GetWeeklyProduct(start, end string) ([]ProductRecord, error) {
	var records []ProductRecord
	totalLimit = 500
	fmt.Printf("\nShowing weekly top products with sales larger than - %d \n\n", totalLimit)

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
  *
FROM (
    SELECT
            oop.Date,
            COALESCE(pc.category_name, ""),
            oop.product_id,
            p.product_name,
            SUM(oop.quantity) quantiy,
            SUM(oop.total) lumpsum
    FROM
            (
        SELECT
            OrderDate as Date,
            product_id,
            SUM(quantity) as quantity,
            SUM(total) as total
        FROM
        (
            SELECT
                %s,
                product_id,
                quantity,
                total
            FROM
                order_product
        ) as op
        GROUP BY
            OrderDate, product_id
        ) as oop

    LEFT OUTER JOIN

    (
        SELECT
               product_id,
               name as product_name
        FROM
               product
    ) as p

    ON oop.product_id = p.product_id

    LEFT OUTER JOIN

    (SELECT
        pc.product_id,
        pc.category_id,
        c.name as category_name
    FROM product_to_category as pc
        INNER JOIN category as c
          on pc.category_id = c.category_id
    WHERE c.language_id = 1) as pc

    ON pc.product_id = oop.product_id
    WHERE oop.Date >= '%s' and oop.Date <= '%s'
    GROUP BY oop.Date, pc.category_name, oop.product_id, p.product_name
    ORDER BY oop.Date desc, total DESC, pc.category_name, oop.product_id
) as final
WHERE lumpsum >= %d
    `
	query := fmt.Sprintf(queryF,
		"CAST(SUBDATE(created_date, WEEKDAY(created_date)) AS DATE) AS OrderDate",
		start,
		end,
		totalLimit)

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

// GetMonthlyProduct returns the latest product sales record for the past n months
func GetMonthlyProduct(start, end string) ([]ProductRecord, error) {
	var records []ProductRecord
	totalLimit = 1000
	fmt.Printf("\nShowing monthly top products with sales larger than - %d \n\n", totalLimit)

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
SELECT *
FROM (
    SELECT
            oop.Date,
            COALESCE(pc.category_name, ""),
            oop.product_id,
            p.product_name,
            SUM(oop.quantity) quantiy,
            SUM(oop.total) lumpsum
    FROM
            (
        SELECT
            OrderDate as Date,
            product_id,
            SUM(quantity) as quantity,
            SUM(total) as total
        FROM
        (
            SELECT
                %s,
                product_id,
                quantity,
                total
            FROM
                order_product
        ) as op
        GROUP BY
            OrderDate, product_id
        ) as oop
    LEFT OUTER JOIN

    (
        SELECT
               product_id,
               name as product_name
        FROM
               product
    ) as p

    ON oop.product_id = p.product_id

    LEFT OUTER JOIN

    (SELECT
        pc.product_id,
        pc.category_id,
        c.name as category_name
    FROM product_to_category as pc
        INNER JOIN category as c
          on pc.category_id = c.category_id
    WHERE c.language_id = 1) as pc

    ON pc.product_id = oop.product_id
    WHERE oop.Date >= '%s' and oop.Date <= '%s'
    GROUP BY oop.Date, pc.category_name, oop.product_id, p.product_name
    ORDER BY oop.Date desc, total DESC, pc.category_name, oop.product_id
) as final
where lumpsum >= %d
    `
	query := fmt.Sprintf(queryF,
		"CAST(DATE_FORMAT(created_date,'%Y-%m-01') as DATE) AS OrderDate",
		start,
		end,
		totalLimit)
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

// GetProductDetails gets product details with the provided product id
func GetProductDetails(id int) ([]Product, error) {
	var products []Product

	db, err := database.GetConnection()
	if err != nil {
		return products, err
	}

	queryF := `
    SELECT
        product_id, code, url, name, price, tb_list_price, discount, is_active
    FROM
        product
    where product_id = %d
`
	query := fmt.Sprintf(queryF, id)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return products, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var rec Product
		err = results.Scan(&rec.ID, &rec.Code, &rec.URL, &rec.Name, &rec.Price, &rec.ListedPrice, &rec.Discount, &rec.Active)
		if err != nil {
			return products, err
		}
		products = append(products, rec) // One record only anyway
	}
	return products, nil
}

// GetProductSalesRecords gets the historical record (weekly) of a particular product
func GetProductSalesRecords(id int) ([]ProductRecord, error) {
	var records []ProductRecord

	db, err := database.GetConnection()
	if err != nil {
		return records, err
	}

	queryF := `
SELECT
	Date,
	product_id,
	SUM(quantity) as quantity,
	SUM(total) as total
FROM
	(
		SELECT
			CAST(SUBDATE(created_date, WEEKDAY(created_date)) AS DATE) AS DATE,
			product_id,
			quantity,
			total
		FROM
			order_product
		WHERE
			product_id = %d
	) as op
GROUP BY
	DATE,
	product_id
ORDER BY
	DATE DESC
`
	query := fmt.Sprintf(queryF, id)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return records, errors.Wrap(err, "cant execute query")
	}

	for results.Next() {
		var rec ProductRecord
		err = results.Scan(&rec.Date, &rec.ProductID, &rec.Quantity, &rec.Total)
		if err != nil {
			return records, err
		}
		records = append(records, rec) // One record only anyway
	}

	return records, nil
}

func GetTopProducts(start string, nrecords int) (map[int][]MonthlySales, error) {
	m := make(map[int][]MonthlySales)
	db, err := database.GetConnection()
	if err != nil {
		return m, err
	}

	var where string // where clause for query
	if start != "" {
		where = fmt.Sprintf(" WHERE PRODUCT_ID is not null AND created_date >= \"%s\" ", start)
	} else {
		where = " WHERE PRODUCT_ID is not null "
	}

	queryF := `
SELECT
	RANK,
    op.PRODUCT_ID,
    p.name as PRODUCT_NAME,
	GrandTotal,
	YearMonth,
	MonthTotal
FROM
(
	SELECT
		%s,
		PRODUCT_ID,
		SUM(Total) as MonthTotal
	FROM
			(SELECT
					%s,
					product_id,
					Total
			FROM order_product
            %s
            ) as day
	GROUP BY DATE, product_id
) as op

INNER JOIN

(SELECT @rn:=@rn+1 AS rank, product_id, GrandTotal
FROM (
  SELECT
		PRODUCT_ID,
		sum(total) as GrandTotal
	FROM
		order_product
	%s

	GROUP BY
		PRODUCT_ID
	ORDER BY
		GrandTotal desc
) t1, (SELECT @rn:=0) t2) as tt

	on op.PRODUCT_ID = tt.PRODUCT_ID

INNER JOIN
    product as p
       ON p.product_id = op.product_id

WHERE tt.rank <= %d
order by rank, YearMonth desc, MonthTotal desc
    `
	query := fmt.Sprintf(queryF,
		"DATE_FORMAT(DATE,'%Y-%m') as YearMonth",
		"CAST(DATE_FORMAT(order_product.created_date,'%Y-%m-01') as DATE) as DATE",
		where,
		where,
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
		err = results.Scan(&rank, &rec.Field, &rec.Name, &rec.GrandTotal, &rec.Month, &rec.Total)
		if v, ok := m[rank]; ok {
			m[rank] = append(v, rec)
		} else {
			m[rank] = []MonthlySales{rec}
		}
	}
	return m, nil
}
