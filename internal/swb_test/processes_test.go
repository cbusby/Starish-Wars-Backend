package swb_test

import (
	"encoding/json"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/Starish-Wars-Backend/internal/swb"
	"github.com/Starish-Wars-Backend/internal/swb/persistence"
)

var _ = ginkgo.Describe("Processes", func() {
	ginkgo.Describe("POST", func() {
		var (
			gameID      string
			game        swb.Game
			err         error
			emptyPlayer swb.Player
		)

		ginkgo.BeforeEach(func() {
			mockPersister := persistence.MockPersister{}
			var g string
			gameID, g, err = swb.Create(mockPersister)
			game = swb.Game{}
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
			gomega.Expect(game.Status).To(gomega.Equal(swb.AWAITING_SHIPS))
		})
	})
})
