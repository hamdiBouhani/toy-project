package storage

import "github.com/pkg/errors"

var (
	// ErrNotFound is the error returned by storages if a resource cannot be found.
	ErrNotFound = errors.New("not found")
)

// Storage is the storage interface used by the server. Implementations are
// required to be able to perform atomic compare-and-swap updates and support standardize on UTC.
type Storage interface {
	Close() error
	Version() (string, error)
}
