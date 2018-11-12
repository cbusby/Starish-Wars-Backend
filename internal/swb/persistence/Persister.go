package persistence

// Persister saves data to a data store
type Persister interface {
	Save(name string, contents string) error

	Read(name string) (string, error)
}
