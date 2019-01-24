package swb

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Validation of initial ship placement", func() {
	var (
		valid   Grid
		invalid Grid
	)

	ginkgo.BeforeEach(func() {
		valid = Grid{}
		valid.Carrier[0] = Coordinate{'A', 1}
		valid.Carrier[1] = Coordinate{'A', 2}
		valid.Carrier[2] = Coordinate{'A', 3}
		valid.Carrier[3] = Coordinate{'A', 4}
		valid.Carrier[4] = Coordinate{'A', 5}

		valid.Battleship[0] = Coordinate{'A', 6}
		valid.Battleship[1] = Coordinate{'A', 7}
		valid.Battleship[2] = Coordinate{'A', 8}
		valid.Battleship[3] = Coordinate{'A', 9}

		valid.Cruiser[0] = Coordinate{'B', 1}
		valid.Cruiser[1] = Coordinate{'B', 2}
		valid.Cruiser[2] = Coordinate{'B', 3}

		valid.Submarine[0] = Coordinate{'B', 4}
		valid.Submarine[1] = Coordinate{'B', 5}
		valid.Submarine[2] = Coordinate{'B', 6}

		valid.Destroyer[0] = Coordinate{'B', 7}
		valid.Destroyer[1] = Coordinate{'B', 8}
	})

	ginkgo.Describe("All ships should be present", func() {
		ginkgo.It("Valid grid should have all ships present", func() {
			gomega.Expect(allShipsOnBoard(valid)).To(gomega.BeTrue())
		})

		ginkgo.It("Grid that is missing a ship is flagged as missing", func() {
			invalid = Grid{}
			invalid.Carrier[0] = Coordinate{'A', 1}
			invalid.Carrier[1] = Coordinate{'A', 2}
			invalid.Carrier[2] = Coordinate{'A', 3}
			invalid.Carrier[3] = Coordinate{'A', 4}
			invalid.Carrier[4] = Coordinate{'A', 5}
			gomega.Expect(allShipsOnBoard(invalid)).NotTo(gomega.BeTrue())
		})
	})
})
