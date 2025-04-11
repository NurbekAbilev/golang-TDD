package conns

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitSQLiteConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
