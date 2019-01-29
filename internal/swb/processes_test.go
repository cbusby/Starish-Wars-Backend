package swb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/cbusby/Starish-Wars-Backend/internal/swb/model"
	"github.com/cbusby/Starish-Wars-Backend/internal/swb/persistence"
	v "github.com/cbusby/Starish-Wars-Backend/internal/swb/validation"
)

var _ = ginkgo.Describe("Processes", func() {
	ginkgo.Describe("POST", func() {
		var (
			gameID      string
			game        model.Game
			err         error
			emptyPlayer model.Player
		)

		ginkgo.BeforeEach(func() {
			mockPersister := persistence.MockPersister{}
			var g string
			gameID, g, err = Create(mockPersister)
			game = model.Game{}
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
			gomega.Expect(game.Status).To(gomega.Equal(model.AWAITING_SHIPS))
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
			persister   persistence.Persister
			contents    string
			err         error
			examplesDir string
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
			pwd, _ := os.Getwd()
			projectDir := filepath.Dir(filepath.Dir(pwd))
			examplesDir = filepath.Join(projectDir, "examples")
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

		ginkgo.Describe("AWAITING_SHIPS", func() {

			ginkgo.It("should update player 1's ship positions when provided", func() {
				newGameBB, _ := ioutil.ReadFile(filepath.Join(examplesDir, "place_ships_request_collapsed.json"))
				newGameString := string(newGameBB)
				contents, _ = Update(persister, gameID, newGameString)
				gomega.Expect(contents).To(gomega.Equal(newGameString))
			})

			ginkgo.It("should update player 2's ship positions and game status when player 1's positions are already provided", func() {
				oldGameStateBB, _ := ioutil.ReadFile(filepath.Join(examplesDir, "place_ships_request_collapsed.json"))
				oldGameString := string(oldGameStateBB)
				persister = persistence.MockPersister{ExpectedGameID: gameID, SavedGameState: oldGameString}
				newGameBB, _ := ioutil.ReadFile(filepath.Join(examplesDir, "player_2_place_ships_request.json"))
				newGameString := string(newGameBB)
				contents, _ := Update(persister, gameID, newGameString)
				var newGameState model.Game
				json.Unmarshal([]byte(contents), &newGameState)
				gomega.Expect(v.ValidateShipPlacement(newGameState.Player1.Ships)).To(gomega.BeTrue())
				gomega.Expect(v.ValidateShipPlacement(newGameState.Player2.Ships)).To(gomega.BeTrue())
				gomega.Expect(newGameState.Status).To(gomega.Equal(model.PLAYER_1_ACTIVE))
			})

			ginkgo.It("should not change game state or status if the provided position is invalid", func() {
				oldGameStateBB, _ := ioutil.ReadFile(filepath.Join(examplesDir, "place_ships_request_collapsed.json"))
				oldGameString := string(oldGameStateBB)
				persister = persistence.MockPersister{ExpectedGameID: gameID, SavedGameState: oldGameString}
				newGame := model.Game{Status: model.AWAITING_SHIPS}
				newGameBB, _ := json.Marshal(newGame)
				newGameString := string(newGameBB)
				contents, _ := Update(persister, gameID, newGameString)
				gomega.Expect(contents).To(gomega.Equal(oldGameString))
			})
		})

		ginkgo.Describe("PLAYER_X_ACTIVE", func() {

			var (
				prev model.Game
				next model.Game
			)

			ginkgo.BeforeEach(func() {
				grid := model.Grid{}
				grid.Carrier[0] = model.Coordinate{"A", 1}
				grid.Carrier[1] = model.Coordinate{"A", 2}
				grid.Carrier[2] = model.Coordinate{"A", 3}
				grid.Carrier[3] = model.Coordinate{"A", 4}
				grid.Carrier[4] = model.Coordinate{"A", 5}

				grid.Battleship[0] = model.Coordinate{"A", 6}
				grid.Battleship[1] = model.Coordinate{"A", 7}
				grid.Battleship[2] = model.Coordinate{"A", 8}
				grid.Battleship[3] = model.Coordinate{"A", 9}

				grid.Cruiser[0] = model.Coordinate{"B", 1}
				grid.Cruiser[1] = model.Coordinate{"B", 2}
				grid.Cruiser[2] = model.Coordinate{"B", 3}

				grid.Submarine[0] = model.Coordinate{"B", 4}
				grid.Submarine[1] = model.Coordinate{"B", 5}
				grid.Submarine[2] = model.Coordinate{"B", 6}

				grid.Destroyer[0] = model.Coordinate{"B", 7}
				grid.Destroyer[1] = model.Coordinate{"B", 8}

				shots := []model.Coordinate{
					model.Coordinate{"A", 1},
				}
				moreShots := []model.Coordinate{
					shots[0],
					model.Coordinate{"A", 2},
				}

				prev = model.Game{}
				prev.Player1 = model.Player{Ships: grid, Shots: shots}
				prev.Player2 = model.Player{Ships: grid, Shots: shots}
				prev.Status = model.PLAYER_2_ACTIVE

				next = prev
				next.Player2.Shots = moreShots
				next.Status = model.PLAYER_1_ACTIVE
			})

			ginkgo.It("returns new game if validations succeed", func() {
				oldGameString, _ := marshalGame(prev)
				persister = persistence.MockPersister{ExpectedGameID: gameID, SavedGameState: oldGameString}
				newGameString, _ := marshalGame(next)
				returnedGameString, _ := Update(persister, gameID, newGameString)
				gomega.Expect(returnedGameString).To(gomega.Equal(newGameString))
			})

			ginkgo.It("returns old game if validations fail", func() {
				oldGameString, _ := marshalGame(prev)
				persister = persistence.MockPersister{ExpectedGameID: gameID, SavedGameState: oldGameString}
				next.Player2.Ships.Destroyer[1] = model.Coordinate{"J", 10}
				newGameString, _ := marshalGame(next)
				returnedGameString, _ := Update(persister, gameID, newGameString)
				gomega.Expect(returnedGameString).To(gomega.Equal(oldGameString))
			})

			ginkgo.It("sets game status to GAME_OVER if all ships are sunk", func() {
				next.Player2.Shots = append([]model.Coordinate(nil), prev.Player1.Ships.Carrier[:]...)
				next.Player2.Shots = append(next.Player2.Shots, prev.Player1.Ships.Battleship[:]...)
				next.Player2.Shots = append(next.Player2.Shots, prev.Player1.Ships.Cruiser[:]...)
				next.Player2.Shots = append(next.Player2.Shots, prev.Player1.Ships.Submarine[:]...)
				next.Player2.Shots = append(next.Player2.Shots, prev.Player1.Ships.Destroyer[:]...)
				prev.Player2.Shots = next.Player2.Shots[:len(next.Player2.Shots)-1]
				oldGameString, _ := marshalGame(prev)
				persister = persistence.MockPersister{ExpectedGameID: gameID, SavedGameState: oldGameString}
				newGameString, _ := marshalGame(next)
				returnedGameString, _ := Update(persister, gameID, newGameString)
				returnedGame, _ := unmarshalGame(returnedGameString)
				gomega.Expect(returnedGame.Status).To(gomega.Equal(model.GAME_OVER))
			})
		})

		ginkgo.Describe("GAME_OVER", func() {

			ginkgo.It("returns final game state when game is over", func() {
				game := model.Game{Status: model.GAME_OVER}
				oldGameString, _ := marshalGame(game)
				persister = persistence.MockPersister{ExpectedGameID: gameID, SavedGameState: oldGameString}
				game.Player1.Shots = append(game.Player1.Shots, model.Coordinate{"A", 1})
				newGameString, _ := marshalGame(game)
				returnedGameString, _ := Update(persister, gameID, newGameString)
				gomega.Expect(returnedGameString).To(gomega.Equal(oldGameString))
			})
		})
	})
})
