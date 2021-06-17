package pg

import (
	"database/sql"
	"errors"
	"log"

	"github.com/cenkalti/backoff/v4"
)

const (
	migrateRetries = 12
)

// Migrate runs migrations. It's supposed to run in production,
// so assume the target database is CockroachDB.
func Migrate(dsn, sqlPath string) DBSetupOpt {
	return func(_ *sql.DB) error {
		if dsn == "" {
			return errors.New("no dsn specified")
		}

		// This is needed because the migrate tool recognizes 'postgres/postgresql'
		// driver part of the DSN and applies the Postgres driver instead of the Cockroach one.
		// dsn = strings.Replace(dsn, "postgresql", "cockroach", 1)
		// dsn = strings.Replace(dsn, "postgres", "cockroach", 1)

		var retry int

		migrateFunc := func() error {
			err := migrateUp(dsn, sqlPath)
			if err != nil {
				log.Println("Retrying to run migrations, err: %v, retries: %v, sqlPath: %v", err, migrateRetries-retry, sqlPath)
				retry++
			}

			return err
		}

		return backoff.Retry(
			migrateFunc,
			backoff.WithMaxRetries(backoff.NewExponentialBackOff(),
				migrateRetries,
			))
	}
}

// migrate use golang-migrate to migrate sql file to databases
func migrateUp(dsn, sqlPath string) error {

	return nil
}
