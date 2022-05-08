//go:build dev
// +build dev

package cmd

import (
	"context"
	"path/filepath"

	"log"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"

	"github.com/spf13/cobra"
)

var migrationName string
var projectDir string

// genMigrateCmd represents the migrate command
var genMigrateCmd = &cobra.Command{
	Use:   "generate-migration",
	Short: "creates migration files",
	Run:   generateMigration,
}

func generateMigration(cmd *cobra.Command, args []string) {

	entSchema := filepath.Join(projectDir, "ent", "schema")

	dsn := (*Conf).Database.GetDsn()
	// Load the graph.
	graph, err := entc.LoadGraph(entSchema, &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	tbls, err := graph.Tables()
	if err != nil {
		log.Fatalln(err)
	}
	// Create a local migration directory.
	d, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalln(err)
	}
	// Open connection to the database.
	dlct, err := sql.Open(Conf.Database.Type, dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// Inspect it and compare it with the graph.
	m, err := schema.NewMigrate(dlct,
		schema.WithDir(d),
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
		schema.WithForeignKeys(true),
	)
	if err != nil {
		log.Fatalln(err)
	}

	if err := m.NamedDiff(context.Background(), migrationName, tbls...); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	rootCmd.AddCommand(genMigrateCmd)

	mflags := genMigrateCmd.Flags()
	mflags.StringVarP(&migrationName, "name", "n", "", "migration name")
	mflags.StringVarP(&projectDir, "project", "p", "", "project directory")

	cobra.MarkFlagRequired(mflags, "name")
	cobra.MarkFlagRequired(mflags, "project")
}
