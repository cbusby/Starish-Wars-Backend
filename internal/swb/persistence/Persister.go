package persistence

// Persister saves data to a data store
type Persister interface {
	Save(gameID string, contents string) error

	Read(gameID string) (string, error)
}
