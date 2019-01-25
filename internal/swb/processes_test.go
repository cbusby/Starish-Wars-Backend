package swb

import (
	"encoding/json"
	"fmt"

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

		ginkgo.It("Should read correctly", func() {
			contents, err = Read(persister, gameID)
			gomega.Expect(contents).To(gomega.Equal(fmt.Sprintf("Read '%s'", gameID)))
		})
	})
})
