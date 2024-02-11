package main

import (
	"auths-services/cmd/api"
	"auths-services/util"
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Println(config.DBSource, err)
		log.Fatal().Msg("Cannot connect to db:")
	}
	fmt.Println(config.DBSource)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Msg(err.Error())
		log.Fatal().Msg("Cannot connect to db:")
	}

	/// run db migration
	runDBMigration(config.MigrationURL, config.DBSource)
	runBarfServer(config, conn)

}

func runDBMigration(migrationURL string, dbSource string) {
	fmt.Println(migrationURL, dbSource)
	migration, err := migrate.New(migrationURL, dbSource)
	fmt.Println(err, "hererer")
	if err != nil {
		log.Fatal().Msgf("cannot create new migrate instance: %w", err)
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Msg("failed to run migration up:")
	}
	log.Info().Msg("db migrated successfully")
}

func runBarfServer(config util.Config, conn *sql.DB) {
	server := api.NewServer(conn, config)
	if server == nil {
		log.Fatal().Msg("cannot create server")
	}
	server.Serve()

}
