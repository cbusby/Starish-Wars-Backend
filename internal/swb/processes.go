package swb

import (
	"encoding/json"
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
	oldGameString, readOldGameErr := persister.Read(gameID)
	if readOldGameErr != nil {
		return "", readOldGameErr
	}
	var oldGame Game
	oldGameUnmarshalErr := json.Unmarshal([]byte(oldGameString), &oldGame)
	if oldGameUnmarshalErr != nil || oldGame.Status == "" {
		return "", fmt.Errorf("Error unmarshaling previous game state")
	}
	var newGame Game
	newGameUnmarshalErr := json.Unmarshal([]byte(requestedGameState), &newGame)
	if newGameUnmarshalErr != nil || newGame.Status == "" {
		return "", fmt.Errorf("Error unmarshaling new game state")
	}
	var updatedGame = oldGame
	if oldGame.Status == AWAITING_SHIPS {
		if !allShipsPresent(oldGame.Player1.Ships) && validateShipPlacement(newGame.Player1.Ships) {
			updatedGame.Player1.Ships = newGame.Player1.Ships
		}
		if !allShipsPresent(oldGame.Player2.Ships) && validateShipPlacement(newGame.Player2.Ships) {
			updatedGame.Player2.Ships = newGame.Player2.Ships
		}
		if allShipsPresent(updatedGame.Player1.Ships) && allShipsPresent(updatedGame.Player2.Ships) {
			updatedGame.Status = PLAYER_1_ACTIVE
		}
	}
	updatedGameByteArray, marshalErr := json.Marshal(updatedGame)
	if marshalErr != nil {
		return "", marshalErr
	}
	updatedGameString := string(updatedGameByteArray)
	saveErr := persister.Save(gameID, updatedGameString)
	if saveErr != nil {
		return "", saveErr
	}
	return updatedGameString, nil
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
