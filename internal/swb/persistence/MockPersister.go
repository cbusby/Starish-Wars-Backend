package persistence

// MockPersister do-nothing implementation of Persister
type MockPersister struct {
	ExpectedGameID string
}

// Save keep function arguments and do nothing
func (m MockPersister) Save(gameID string, contents string) error {
	return nil
}

func (m MockPersister) Read(gameID string) (string, error) {
	return "Read '" + gameID + "'", nil
}
