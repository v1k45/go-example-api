package cmd

import (
	"log"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1k45/shitpost/db"
)

var migrator *migrate.Migrate

var migrationsCmd = &cobra.Command{
	Use:   "migrations",
	Short: "Run database migrations",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		migrator, err = db.Migrate(viper.GetString("database_url"))
		if err != nil {
			log.Fatalf("Error initializing migrator: %v", err)
		}
	},
}

func init() {
	migrationsCmd.AddCommand(migrateUpCmd, migrateDownCmd, migrateToCmd, migrateDrop)
	rootCmd.AddCommand(migrationsCmd)
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Migrate the database up to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Error migrating up: %v", err)
		}

		log.Println("Migrated up to latest version")
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Migrate the database down to the previous version",
	Run: func(cmd *cobra.Command, args []string) {
		if err := migrator.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Error migrating down: %v", err)
		}

		log.Println("Migrated down to previous version")
	},
}

var migrateToCmd = &cobra.Command{
	Use:   "to <version>",
	Short: "Migrate the database to a specific version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		migrateToVersion, err := strconv.Atoi(args[0])
		if err != nil || migrateToVersion < 0 {
			log.Fatalf("Invalid migration version: %v", err)
		}

		if migrateToVersion == 0 {
			log.Fatal("Missing migration version", migrateToVersion)
		}

		if err := migrator.Migrate(uint(migrateToVersion)); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Error migrating to version %d: %v", migrateToVersion, err)
		}
		log.Printf("Migrated to version %d", migrateToVersion)
	},
}

var migrateDrop = &cobra.Command{
	Use:   "drop",
	Short: "Drop all tables from the database",
	Run: func(cmd *cobra.Command, args []string) {
		if err := migrator.Drop(); err != nil {
			log.Fatalf("Error dropping tables: %v", err)
		}

		log.Println("Dropped all tables")
	},
}
