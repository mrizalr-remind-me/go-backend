package sqlx

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func New() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
