package persistence

import (
	"fmt"
)

// MockPersister do-nothing implementation of Persister
type MockPersister struct {
	ExpectedGameID string
	SavedGameState string
}

// Save keep function arguments and do nothing
func (m MockPersister) Save(gameID string, contents string) error {
	return nil
}

// Read return game state the mock was constructed with, an informational message, or an error
func (m MockPersister) Read(gameID string) (string, error) {
	if gameID == m.ExpectedGameID {
		if len(m.SavedGameState) == 0 {
			return "Read '" + gameID + "'", nil
		}
		return m.SavedGameState, nil
	}
	return "", fmt.Errorf("Not found: %s", gameID)
}
