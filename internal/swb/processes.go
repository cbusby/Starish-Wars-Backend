package swb

import (
	"fmt"
	"log"
	"os"

	"github.com/rs/xid"

	"github.com/cbusby/Starish-Wars-Backend/internal/swb/persistence"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

// Create creates a new game, writes the state to S3, and returns the GameID and the state
func Create(persister persistence.Persister) (string, string, error) {
	gameID := xid.New().String()
	body := newGame()

	uploadErr := persister.Save(gameID, body)
	if uploadErr != nil {
		return "", "", uploadErr
	}

	return gameID, body, nil
}

// Update checks that the new state is valid, including comparing the previous game state to the new one, and updates the new game state accordingly
func Update(persister persistence.Persister, gameID string, requestedGameState string) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

// Read reads the current state of an existing game
func Read(persister persistence.Persister, gameID string) (string, error) {
	return persister.Read(gameID)
}

func newGame() string {
	return `{
	"status": "AWAITING_SHIPS",
	"player_1": {},
	"player_2": {}
}`
}
