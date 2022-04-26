package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DB() *sql.DB {
	db, err := sql.Open("mysql", "root:skymkey08@/Todo")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
	