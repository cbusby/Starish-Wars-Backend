package swb

import (
	"testing"

	"github.com/Starish-Wars-Backend/internal/swb/persistence"
)

// TestCreate test the Create function
func TestCreate(t *testing.T) {
	mockPersister := persistence.MockPersister{}
	gameID, response, err := Create(mockPersister)
	if err != nil {
		t.Errorf("Did not expect an error")
	}
	if gameID == "" {
		t.Errorf("gameID should not be empty")
	}
	if response == "" {
		t.Errorf("response should not be empty")
	}
}
