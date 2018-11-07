package persistence

// MockPersister do-nothing implementation of Persister
type MockPersister struct {
}

var savedName string
var savedContents string

// Save keep function arguments and do nothing
func (m MockPersister) Save(name string, contents string) error {
	savedName = name
	savedContents = contents
	return nil
}
