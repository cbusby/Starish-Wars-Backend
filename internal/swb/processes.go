package swb

import (
	"log"
	"os"

	"github.com/Starish-Wars-Backend/internal/swb/persistence"
	"github.com/rs/xid"
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

func newGame() string {
	return `{
	"status": "AWAITING_SHIPS",
	"player_1": {},
	"player_2": {}
}`
}
