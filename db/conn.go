package db

import (
	"database/sql"
	"strings"
)

func Open(databaseUrl string) (*sql.DB, error) {
	return sql.Open("sqlite3", strings.ReplaceAll(databaseUrl, "sqlite3://", ""))
}
