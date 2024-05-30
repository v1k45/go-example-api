package cmd

import (
	"log"
	"log/slog"
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
			slog.Error("migrations_initialization_failed", "error", err)
			log.Fatalf("Error initializing migrations: %v", err)
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
			slog.Error("migrations_up_failed", "error", err)
			log.Fatalf("Error migrating up: %v", err)
		}

		slog.Info("migrations_up_succeeded")
		log.Println("Migrated up to latest version")
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Migrate the database down to the previous version",
	Run: func(cmd *cobra.Command, args []string) {
		if err := migrator.Down(); err != nil && err != migrate.ErrNoChange {
			slog.Error("migrations_down_failed", "error", err)
			log.Fatalf("Error migrating down: %v", err)
		}

		slog.Info("migrations_down_succeeded")
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
			slog.Error("migrations_invalid_version", "error", err)
			log.Fatalf("Invalid migration version: %v", err)
		}

		if migrateToVersion == 0 {
			slog.Error("migrations_missing_version", "version", migrateToVersion)
			log.Fatal("Missing migration version", migrateToVersion)
		}

		if err := migrator.Migrate(uint(migrateToVersion)); err != nil && err != migrate.ErrNoChange {
			slog.Error("migrations_to_failed", "version", migrateToVersion, "error", err)
			log.Fatalf("Error migrating to version %d: %v", migrateToVersion, err)
		}

		slog.Info("migrations_to_succeeded", "version", migrateToVersion)
		log.Printf("Migrated to version %d", migrateToVersion)
	},
}

var migrateDrop = &cobra.Command{
	Use:   "drop",
	Short: "Drop all tables from the database",
	Run: func(cmd *cobra.Command, args []string) {
		if err := migrator.Drop(); err != nil {
			slog.Error("migrations_drop_failed", "error", err)
			log.Fatalf("Error dropping tables: %v", err)
		}

		log.Println("Dropped all tables")
	},
}
