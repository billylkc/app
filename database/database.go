package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// GetConnection return the connection of the mysql database
func GetConnection() (*sql.DB, error) {
	var db *sql.DB
	secret := os.Getenv("JMALL_CONNECT")
	if secret == "" {
		return db, fmt.Errorf("missing environment variable JMALL_CONNECT. Please check.")
	}
	db, err := sql.Open("mysql", secret)
	if err != nil {
		return db, errors.Wrap(err, "unable to connect to mysql")
	}
	return db, err
}
