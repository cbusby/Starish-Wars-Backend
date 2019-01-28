package swb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/cbusby/Starish-Wars-Backend/internal/swb/persistence"
)

var _ = ginkgo.Describe("Processes", func() {
	ginkgo.Describe("POST", func() {
		var (
			gameID      string
			game        Game
			err         error
			emptyPlayer Player
		)

		ginkgo.BeforeEach(func() {
			mockPersister := persistence.MockPersister{}
			var g string
			gameID, g, err = Create(mockPersister)
			game = Game{}
			json.Unmarshal([]byte(g), &game)
		})

		ginkgo.It("Should not return an error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("Should return a nonempty gameID", func() {
			gomega.Expect(gameID).NotTo(gomega.Equal(""))
		})

		ginkgo.It("should return empty object for player 1", func() {
			gomega.Expect(game.Player1).To(gomega.Equal(emptyPlayer))
		})

		ginkgo.It("should return empty object for player 2", func() {
			gomega.Expect(game.Player2).To(gomega.Equal(emptyPlayer))
		})

		ginkgo.It("should return status AWAITING_SHIPS", func() {
			gomega.Expect(game.Status).To(gomega.Equal(AWAITING_SHIPS))
		})
	})

	ginkgo.Describe("GET", func() {

		gameID := "123abc"
		var (
			persister persistence.Persister
			contents  string
			err       error
		)

		ginkgo.BeforeEach(func() {
			persister = persistence.MockPersister{ExpectedGameID: gameID}
		})

		ginkgo.It("Should not return an error if valid gameID is given", func() {
			contents, err = Read(persister, gameID)
			gomega.Expect(err).To(gomega.BeNil())
		})

		ginkgo.It("Should return an error if an invalid gameID is given", func() {
			contents, err = Read(persister, "hoohah")
			gomega.Expect(err).NotTo(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("Not found: hoohah"))
		})

		ginkgo.It("Should read correctly", func() {
			contents, err = Read(persister, gameID)
			gomega.Expect(contents).To(gomega.Equal(fmt.Sprintf("Read '%s'", gameID)))
		})
	})

	ginkgo.Describe("PUT", func() {

		gameID := "123abc"
		var (
			persister persistence.Persister
			contents  string
			err       error
		)

		ginkgo.BeforeEach(func() {
			persister = persistence.MockPersister{
				ExpectedGameID: gameID,
				SavedGameState: `{
	"status": "AWAITING_SHIPS",
	"player_1": {},
	"player_2": {}
}`,
			}
		})

		ginkgo.It("should return an error if an invalid game id is given", func() {
			_, err = Update(persister, "hoohah", "{}")
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("should return an error if saved game state is invalid", func() {
			persister = persistence.MockPersister{ExpectedGameID: gameID, SavedGameState: `{"hoo": "hah"}`}
			_, err = Update(persister, gameID, "{}")
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("should return an error if new game state is invalid", func() {
			_, err = Update(persister, gameID, `{"hoo": "hah"}`)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("should update player 1's ship positions when provided", func() {
			pwd, _ := os.Getwd()
			projectDir := filepath.Dir(filepath.Dir(pwd))
			newGameBB, _ := ioutil.ReadFile(filepath.Join(projectDir, "examples", "place_ships_request_collapsed.json"))
			newGameString := string(newGameBB)
			contents, _ = Update(persister, gameID, newGameString)
			gomega.Expect(contents).To(gomega.Equal(newGameString))
		})

		ginkgo.It("should update player 2's ship positions when player 1's positions are already provided", func() {
			pwd, _ := os.Getwd()
			projectDir := filepath.Dir(filepath.Dir(pwd))
			oldGameStateBB, _ := ioutil.ReadFile(filepath.Join(projectDir, "examples", "place_ships_request_collapsed.json"))
			oldGameString := string(oldGameStateBB)
			persister = persistence.MockPersister{ExpectedGameID: gameID, SavedGameState: oldGameString}
			newGameBB, _ := ioutil.ReadFile(filepath.Join(projectDir, "examples", "player_2_place_ships_request.json"))
			newGameString := string(newGameBB)
			contents, _ := Update(persister, gameID, newGameString)
			var newGameState Game
			json.Unmarshal([]byte(contents), &newGameState)
			gomega.Expect(validateShipPlacement(newGameState.Player1.Ships)).To(gomega.BeTrue())
			gomega.Expect(validateShipPlacement(newGameState.Player2.Ships)).To(gomega.BeTrue())
			gomega.Expect(newGameState.Status).To(gomega.Equal(PLAYER_1_ACTIVE))
		})
	})
})
