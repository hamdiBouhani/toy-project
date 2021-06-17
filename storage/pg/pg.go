package pg

import (
	"database/sql"
	"fmt"
	"log"
	"toy-project/storage"

	"github.com/cenkalti/backoff/v4"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/xo/dburl"
)

const defaultMaxOpenConns = 1000
const defaultRetries = 6

// Open connect to a database
// if the database not exit then create it first
// automatically migrate the database script after connecting finished
func Open(dsn, migrationPath string) (storage.Storage, error) {

	var db *sqlx.DB
	var err error
	retries := 0

	n := defaultRetries - 1

	u, err := dburl.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse database address, %w", err)
	}
	driver, dataSource := u.Driver, u.DSN
	try := func() error {
		log.Printf("Connecting to db, driver: %v, dsn: %v retries-left: %v", driver, dataSource, n-retries)

		// Check if DSN is in order. If not, return nil and check err value.
		db, err = sqlx.Open(driver, dataSource)
		if err != nil {
			// Bad DSN, we quit immediately
			err = fmt.Errorf("bad dsn, %w", err)
			return nil
		}

		if dbErr := db.Ping(); dbErr != nil {

			retries++
			return fmt.Errorf("couldn't ping db, %w", dbErr)
		}

		return nil
	}

	boff := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), uint64(n))
	errBackoff := backoff.Retry(try, boff)

	// Bad dsn.
	if err != nil {
		return nil, err
	}

	// Couldn't connect after n attempts.
	if errBackoff != nil {
		return nil, errBackoff
	}

	return &conn{
		db: db,
	}, nil
}

// conn is the main database connection.
type conn struct {
	db     *sqlx.DB
	logger logrus.FieldLogger
}

func (c *conn) Close() error {
	return c.db.Close()
}

func (c *conn) Version() (i string, err error) {
	err = c.QueryRow(`
		select
			version
		from schema_migrations limit 1`).Scan(&i)
	if err != nil {
		if err == sql.ErrNoRows {
			return i, storage.ErrNotFound
		}
		return i, errors.Wrap(err, "query schema migration version failed")
	}
	return i, nil
}

// conn implements the same method signatures as encoding/sql.DB.

func (c *conn) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.db.Exec(query, args...)
}

func (c *conn) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.Query(query, args...)
}

func (c *conn) QueryRow(query string, args ...interface{}) *sql.Row {
	return c.db.QueryRow(query, args...)
}

// ExecTx runs a method which operates on a transaction.
func (c *conn) ExecTx(fn func(tx *trans) error) error {
	sqlTx, err := c.db.Beginx()
	if err != nil {
		return err
	}
	t := &trans{
		tx: sqlTx,
	}
	if err := fn(t); err != nil {
		_ = sqlTx.Rollback()
		return err
	}
	return sqlTx.Commit()
}

type trans struct {
	tx *sqlx.Tx
}

// trans implements the same method signatures as encoding/sql.Tx.

func (t *trans) Exec(query string, args ...interface{}) (sql.Result, error) {
	return t.tx.Exec(query, args...)
}

func (t *trans) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return t.tx.Query(query, args...)
}

func (t *trans) QueryRow(query string, args ...interface{}) *sql.Row {
	return t.tx.QueryRow(query, args...)
}
