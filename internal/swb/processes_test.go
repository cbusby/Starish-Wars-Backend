package swb

import (
	"encoding/json"
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
	responseObject := Game{}
	json.Unmarshal([]byte(response), &responseObject)
	if responseObject.Player1 != (Player{}) {
		t.Errorf("Expected player 1 to be nil when game is created")
	}
	if responseObject.Player2 != (Player{}) {
		t.Errorf("Expected player 2 to be nil when game is created")
	}
	if responseObject.Status != "AWAITING_SHIPS" {
		t.Errorf("Expected game status to be AWAITING_SHIPS")
	}
}
