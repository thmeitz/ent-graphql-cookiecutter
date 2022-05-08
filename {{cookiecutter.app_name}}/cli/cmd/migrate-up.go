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

// migrateCmd represents the migrate command
var migrateUpCmd = &cobra.Command{
	Use:   "migrate-up",
	Short: "migrate the database up",
	Run:   migrateUp,
}

var migrateUpDir string

func init() {
	rootCmd.AddCommand(migrateUpCmd)
	mflags := migrateUpCmd.Flags()
	mflags.StringVarP(&migrateUpDir, "migrations", "m", "", "migrations directory")
	cobra.MarkFlagRequired(mflags, "migrations")
}

func migrateUp(cmd *cobra.Command, args []string) {
	var driver database.Driver
	logger := zap.NewExample()

	dsn := (*Conf).Database.GetDsn()
	if Conf.Database.Type == "mysql" {
		dsn = fmt.Sprintf("%v&multiStatements=true", dsn)
	}
	db, err := sql.Open(Conf.Database.Type, dsn)
	if err != nil {
		logger.Fatal("migrate-up", zap.Error(err))
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

	logger.Info("migrate-up", zap.String("dbtype", Conf.Database.Type))

	migrationsDirectory := fmt.Sprintf("file://%v", migrateUpDir)

	logger.Info("migrate-up", zap.Any("migrations", migrationsDirectory))

	m, err := migrate.NewWithDatabaseInstance(
		migrationsDirectory,
		Conf.Database.Type, driver)
	if err != nil {
		logger.Error("migrate-up", zap.Error(err))
	}
	err = m.Up()
	if err != nil {
		logger.Error("migrate-up", zap.Error(err))
	}
}
