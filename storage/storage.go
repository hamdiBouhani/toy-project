package storage

// Storage is the storage interface used by the server. Implementations are
// required to be able to perform atomic compare-and-swap updates and support standardize on UTC.
type Storage interface {
	Close() error
	Version() (string, error)
}
