package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// migrateDownCmd represents the migrateDown command
var migrateDownCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate the database stepwise up or down",
	Run:   migrateUpDown,
}

var migrateDownDir string
var migrateSteps int

func init() {
	rootCmd.AddCommand(migrateDownCmd)
	mflags := migrateDownCmd.Flags()
	mflags.StringVarP(&migrateDownDir, "migrations", "m", "", "migrations directory")
	mflags.IntVarP(&migrateSteps, "steps", "s", 0, "migration steps -1 (migrate 1 step down) +1 (migrate 1 step up)")
	cobra.MarkFlagRequired(mflags, "migrations")
	cobra.MarkFlagRequired(mflags, "steps")
}

func migrateUpDown(cmd *cobra.Command, args []string) {
	var driver database.Driver
	logger := zap.NewExample()

	dsn := (*Conf).Database.GetDsn()
	if Conf.Database.Type == "mysql" {
		dsn = fmt.Sprintf("%v&multiStatements=true", dsn)
	}
	db, err := sql.Open(Conf.Database.Type, dsn)
	if err != nil {
		logger.Fatal("downmigrate", zap.Error(err))
	}
	defer db.Close()

	switch Conf.Database.Type {
	case "mysql":
		driver, err = mysql.WithInstance(db, &mysql.Config{})
	case "postgres":
		driver, err = postgres.WithInstance(db, &postgres.Config{})
	case "sqlite":
		driver, err = sqlite3.WithInstance(db, &sqlite3.Config{})
	}

	logger.Info("migrate-down", zap.String("dbtype", Conf.Database.Type))

	migrationsDirectory := fmt.Sprintf("file://%v", migrateDownDir)

	logger.Info("migrate-down", zap.Any("migrations", migrationsDirectory))

	m, err := migrate.NewWithDatabaseInstance(
		migrationsDirectory,
		Conf.Database.Type, driver)
	if err != nil {
		logger.Error("migrate-down", zap.Error(err))
	}

	err = m.Steps(migrateSteps)
	if err != nil {
		logger.Error("migrate-down", zap.Error(err))
	}
}
