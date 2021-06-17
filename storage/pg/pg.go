package pg

import (
	"database/sql"
	"toy-project/storage"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const defaultMaxOpenConns = 1000

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
