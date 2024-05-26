package db

import (
	"embed"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var fs embed.FS

func Migrate(dbUrl string) (*migrate.Migrate, error) {
	src, err := iofs.New(fs, "migrations")
	if err != nil {
		return nil, err
	}
	return migrate.NewWithSourceInstance("iofs", src, dbUrl)
}
