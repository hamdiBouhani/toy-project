package pg

import (
	"database/sql"

	"go.uber.org/zap"
)

// DBSetupOpt is a function to perform some setup task on a database.
type DBSetupOpt func(*sql.DB) error

// DBSetups set db setups
func DBSetups(setups ...DBSetupOpt) ConnectOption {
	return func(o *connectOptions) {
		o.dbSetups = setups
	}
}

type connectOptions struct {
	logger   *zap.Logger
	retries  int
	dbSetups []DBSetupOpt
}

// ConnectOption is optional config of connecting
type ConnectOption func(*connectOptions)

// Logger set logger
func Logger(logger *zap.Logger) ConnectOption {
	return func(o *connectOptions) {
		o.logger = logger
	}
}

// Retry set retry times
func Retry(n int) ConnectOption {
	return func(o *connectOptions) {
		if n >= 0 {
			o.retries = n
		}
	}
}

func applyOptions(opts ...ConnectOption) *connectOptions {
	o := &connectOptions{
		logger:  zap.NewNop(),
		retries: defaultRetries}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
