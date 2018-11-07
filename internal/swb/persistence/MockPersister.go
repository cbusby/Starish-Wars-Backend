package persistence

// MockPersister do-nothing implementation of Persister
type MockPersister struct {
}

// Save keep function arguments and do nothing
func (m MockPersister) Save(name string, contents string) error {
	return nil
}
