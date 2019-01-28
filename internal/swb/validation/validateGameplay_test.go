package validation

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/cbusby/Starish-Wars-Backend/internal/swb/model"
)

var _ = ginkgo.Describe("Validation of gameplay", func() {
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

	ginkgo.It("successfully detects when ships are the same from move to move", func() {
		gomega.Expect(ShipsSame(prev.Player1, next.Player1)).To(gomega.BeTrue())
	})

	ginkgo.It("successfully detects when inactive player's shot list is the same from move to move", func() {
		gomega.Expect(ShotsSame(prev.Player1, next.Player1)).To(gomega.BeTrue())
	})

	ginkgo.It("successfully detects when active player's shot list is the same with one new shot from move to move", func() {
		gomega.Expect(OneNewShot(prev.Player2, next.Player2)).To(gomega.BeTrue())
	})

	ginkgo.It("correctly detects when all of a player's ships haven't been hit", func() {
		gomega.Expect(AllShipsHit(next.Player1, next.Player2.Shots)).To(gomega.BeFalse())
	})
})
