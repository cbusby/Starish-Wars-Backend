package swb

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/cbusby/Starish-Wars-Backend/internal/swb/model"
)

var _ = ginkgo.Describe("Validation of initial ship placement", func() {
	var (
		valid   model.Grid
		invalid model.Grid
	)

	ginkgo.BeforeEach(func() {
		valid = model.Grid{}
		valid.Carrier[0] = model.Coordinate{"A", 1}
		valid.Carrier[1] = model.Coordinate{"A", 2}
		valid.Carrier[2] = model.Coordinate{"A", 3}
		valid.Carrier[3] = model.Coordinate{"A", 4}
		valid.Carrier[4] = model.Coordinate{"A", 5}

		valid.Battleship[0] = model.Coordinate{"A", 6}
		valid.Battleship[1] = model.Coordinate{"A", 7}
		valid.Battleship[2] = model.Coordinate{"A", 8}
		valid.Battleship[3] = model.Coordinate{"A", 9}

		valid.Cruiser[0] = model.Coordinate{"B", 1}
		valid.Cruiser[1] = model.Coordinate{"B", 2}
		valid.Cruiser[2] = model.Coordinate{"B", 3}

		valid.Submarine[0] = model.Coordinate{"B", 4}
		valid.Submarine[1] = model.Coordinate{"B", 5}
		valid.Submarine[2] = model.Coordinate{"B", 6}

		valid.Destroyer[0] = model.Coordinate{"B", 7}
		valid.Destroyer[1] = model.Coordinate{"B", 8}
	})

	ginkgo.Describe("All ships should be present", func() {
		ginkgo.It("flags a model.Grid as valid if all ships present", func() {
			gomega.Expect(allShipsPresent(valid)).To(gomega.BeTrue())
		})

		ginkgo.It("flags a model.Grid as invalid if a ship is missing", func() {
			invalid = model.Grid{}
			invalid.Carrier[0] = model.Coordinate{"A", 1}
			invalid.Carrier[1] = model.Coordinate{"A", 2}
			invalid.Carrier[2] = model.Coordinate{"A", 3}
			invalid.Carrier[3] = model.Coordinate{"A", 4}
			invalid.Carrier[4] = model.Coordinate{"A", 5}
			gomega.Expect(allShipsPresent(invalid)).NotTo(gomega.BeTrue())
		})
	})

	ginkgo.Describe("All ships should fall on the 10x10 model.Grid", func() {
		ginkgo.It("flags a model.Grid as valid if all ships fall on the model.Grid", func() {
			gomega.Expect(allShipsOnGrid(valid)).To(gomega.BeTrue())
		})

		ginkgo.It("flags a model.Grid as invalid if a ship falls off the left side of the model.Grid", func() {
			invalid = model.Grid{}
			invalid.Carrier[0] = model.Coordinate{"A", 1}
			invalid.Carrier[1] = model.Coordinate{"A", 2}
			invalid.Carrier[2] = model.Coordinate{"A", 3}
			invalid.Carrier[3] = model.Coordinate{"A", 4}
			invalid.Carrier[4] = model.Coordinate{"A", 5}

			invalid.Battleship[0] = model.Coordinate{"A", 6}
			invalid.Battleship[1] = model.Coordinate{"A", 7}
			invalid.Battleship[2] = model.Coordinate{"A", 8}
			invalid.Battleship[3] = model.Coordinate{"A", 9}

			invalid.Cruiser[0] = model.Coordinate{"B", 1}
			invalid.Cruiser[1] = model.Coordinate{"B", 2}
			invalid.Cruiser[2] = model.Coordinate{"B", 3}

			invalid.Submarine[0] = model.Coordinate{"B", 4}
			invalid.Submarine[1] = model.Coordinate{"B", 5}
			invalid.Submarine[2] = model.Coordinate{"B", 6}

			invalid.Destroyer[0] = model.Coordinate{"C", 0}
			invalid.Destroyer[1] = model.Coordinate{"C", 1}
			gomega.Expect(allShipsOnGrid(invalid)).NotTo(gomega.BeTrue())
		})

		ginkgo.It("flags a model.Grid as invalid if a ship falls off the right side of the model.Grid", func() {
			invalid = model.Grid{}
			invalid.Carrier[0] = model.Coordinate{"A", 1}
			invalid.Carrier[1] = model.Coordinate{"A", 2}
			invalid.Carrier[2] = model.Coordinate{"A", 3}
			invalid.Carrier[3] = model.Coordinate{"A", 4}
			invalid.Carrier[4] = model.Coordinate{"A", 5}

			invalid.Battleship[0] = model.Coordinate{"A", 6}
			invalid.Battleship[1] = model.Coordinate{"A", 7}
			invalid.Battleship[2] = model.Coordinate{"A", 8}
			invalid.Battleship[3] = model.Coordinate{"A", 9}

			invalid.Cruiser[0] = model.Coordinate{"B", 1}
			invalid.Cruiser[1] = model.Coordinate{"B", 2}
			invalid.Cruiser[2] = model.Coordinate{"B", 3}

			invalid.Submarine[0] = model.Coordinate{"B", 4}
			invalid.Submarine[1] = model.Coordinate{"B", 5}
			invalid.Submarine[2] = model.Coordinate{"B", 6}

			invalid.Destroyer[0] = model.Coordinate{"A", 10}
			invalid.Destroyer[1] = model.Coordinate{"A", 11}
			gomega.Expect(allShipsOnGrid(invalid)).NotTo(gomega.BeTrue())
		})

		ginkgo.It("flags a model.Grid as invalid if a ship falls off the top side of the model.Grid", func() {
			invalid = model.Grid{}
			invalid.Carrier[0] = model.Coordinate{"A", 1}
			invalid.Carrier[1] = model.Coordinate{"A", 2}
			invalid.Carrier[2] = model.Coordinate{"A", 3}
			invalid.Carrier[3] = model.Coordinate{"A", 4}
			invalid.Carrier[4] = model.Coordinate{"A", 5}

			invalid.Battleship[0] = model.Coordinate{"A", 6}
			invalid.Battleship[1] = model.Coordinate{"A", 7}
			invalid.Battleship[2] = model.Coordinate{"A", 8}
			invalid.Battleship[3] = model.Coordinate{"A", 9}

			invalid.Cruiser[0] = model.Coordinate{"B", 1}
			invalid.Cruiser[1] = model.Coordinate{"B", 2}
			invalid.Cruiser[2] = model.Coordinate{"B", 3}

			invalid.Submarine[0] = model.Coordinate{"B", 4}
			invalid.Submarine[1] = model.Coordinate{"B", 5}
			invalid.Submarine[2] = model.Coordinate{"B", 6}

			invalid.Destroyer[0] = model.Coordinate{"z", 1}
			invalid.Destroyer[1] = model.Coordinate{"A", 1}
			gomega.Expect(allShipsOnGrid(invalid)).NotTo(gomega.BeTrue())
		})

		ginkgo.It("flags a model.Grid as invalid if a ship falls off the bottom side of the model.Grid", func() {
			invalid = model.Grid{}
			invalid.Carrier[0] = model.Coordinate{"A", 1}
			invalid.Carrier[1] = model.Coordinate{"A", 2}
			invalid.Carrier[2] = model.Coordinate{"A", 3}
			invalid.Carrier[3] = model.Coordinate{"A", 4}
			invalid.Carrier[4] = model.Coordinate{"A", 5}

			invalid.Battleship[0] = model.Coordinate{"A", 6}
			invalid.Battleship[1] = model.Coordinate{"A", 7}
			invalid.Battleship[2] = model.Coordinate{"A", 8}
			invalid.Battleship[3] = model.Coordinate{"A", 9}

			invalid.Cruiser[0] = model.Coordinate{"B", 1}
			invalid.Cruiser[1] = model.Coordinate{"B", 2}
			invalid.Cruiser[2] = model.Coordinate{"B", 3}

			invalid.Submarine[0] = model.Coordinate{"B", 4}
			invalid.Submarine[1] = model.Coordinate{"B", 5}
			invalid.Submarine[2] = model.Coordinate{"B", 6}

			invalid.Destroyer[0] = model.Coordinate{"J", 1}
			invalid.Destroyer[1] = model.Coordinate{"K", 1}
			gomega.Expect(allShipsOnGrid(invalid)).NotTo(gomega.BeTrue())
		})
	})

	ginkgo.Describe("All ships should be horizontal or vertical", func() {
		ginkgo.It("flags a model.Grid as valid if all ships are horizontal or vertical", func() {
			gomega.Expect(allShipsHorizontalOrVertical(valid)).To(gomega.BeTrue())
		})

		ginkgo.It("flags a model.Grid as invalid if a ship is diagonal", func() {
			invalid = model.Grid{}
			invalid.Carrier[0] = model.Coordinate{"A", 1}
			invalid.Carrier[1] = model.Coordinate{"A", 2}
			invalid.Carrier[2] = model.Coordinate{"A", 3}
			invalid.Carrier[3] = model.Coordinate{"A", 4}
			invalid.Carrier[4] = model.Coordinate{"A", 5}

			invalid.Battleship[0] = model.Coordinate{"A", 6}
			invalid.Battleship[1] = model.Coordinate{"A", 7}
			invalid.Battleship[2] = model.Coordinate{"A", 8}
			invalid.Battleship[3] = model.Coordinate{"A", 9}

			invalid.Cruiser[0] = model.Coordinate{"B", 1}
			invalid.Cruiser[1] = model.Coordinate{"B", 2}
			invalid.Cruiser[2] = model.Coordinate{"B", 3}

			invalid.Submarine[0] = model.Coordinate{"B", 4}
			invalid.Submarine[1] = model.Coordinate{"B", 5}
			invalid.Submarine[2] = model.Coordinate{"B", 6}

			invalid.Destroyer[0] = model.Coordinate{"B", 7}
			invalid.Destroyer[1] = model.Coordinate{"C", 8}

			gomega.Expect(allShipsHorizontalOrVertical(invalid)).NotTo(gomega.BeTrue())
		})

		ginkgo.It("flags a model.Grid as invalid if a ship is in two columns", func() {
			invalid = model.Grid{}
			invalid.Carrier[0] = model.Coordinate{"A", 1}
			invalid.Carrier[1] = model.Coordinate{"A", 2}
			invalid.Carrier[2] = model.Coordinate{"A", 3}
			invalid.Carrier[3] = model.Coordinate{"A", 4}
			invalid.Carrier[4] = model.Coordinate{"A", 5}

			invalid.Battleship[0] = model.Coordinate{"A", 6}
			invalid.Battleship[1] = model.Coordinate{"A", 7}
			invalid.Battleship[2] = model.Coordinate{"A", 8}
			invalid.Battleship[3] = model.Coordinate{"A", 9}

			invalid.Cruiser[0] = model.Coordinate{"B", 1}
			invalid.Cruiser[1] = model.Coordinate{"B", 2}
			invalid.Cruiser[2] = model.Coordinate{"B", 3}

			invalid.Submarine[0] = model.Coordinate{"B", 4}
			invalid.Submarine[1] = model.Coordinate{"C", 5}
			invalid.Submarine[2] = model.Coordinate{"D", 4}

			invalid.Destroyer[0] = model.Coordinate{"B", 7}
			invalid.Destroyer[1] = model.Coordinate{"B", 8}

			gomega.Expect(allShipsHorizontalOrVertical(invalid)).NotTo(gomega.BeTrue())
		})

		ginkgo.It("flags a model.Grid as invalid if a ship is in two rows", func() {
			invalid = model.Grid{}
			invalid.Carrier[0] = model.Coordinate{"A", 1}
			invalid.Carrier[1] = model.Coordinate{"A", 2}
			invalid.Carrier[2] = model.Coordinate{"A", 3}
			invalid.Carrier[3] = model.Coordinate{"A", 4}
			invalid.Carrier[4] = model.Coordinate{"A", 5}

			invalid.Battleship[0] = model.Coordinate{"A", 6}
			invalid.Battleship[1] = model.Coordinate{"A", 7}
			invalid.Battleship[2] = model.Coordinate{"A", 8}
			invalid.Battleship[3] = model.Coordinate{"A", 9}

			invalid.Cruiser[0] = model.Coordinate{"B", 1}
			invalid.Cruiser[1] = model.Coordinate{"B", 2}
			invalid.Cruiser[2] = model.Coordinate{"B", 3}

			invalid.Submarine[0] = model.Coordinate{"B", 4}
			invalid.Submarine[1] = model.Coordinate{"C", 5}
			invalid.Submarine[2] = model.Coordinate{"B", 6}

			invalid.Destroyer[0] = model.Coordinate{"B", 7}
			invalid.Destroyer[1] = model.Coordinate{"B", 8}

			gomega.Expect(allShipsHorizontalOrVertical(invalid)).NotTo(gomega.BeTrue())
		})
	})

	ginkgo.Describe("All ships are in contiguous spaces", func() {
		ginkgo.It("flags as valid a model.Grid where all ships are in contiguous spaces", func() {
			gomega.Expect(allShipsInTouchingSpaces(valid)).To(gomega.BeTrue())
		})

		ginkgo.It("flags as valid a model.Grid where all ships are in contiguous spaces, even if the model.Coordinates are not in order horizontally", func() {
			outOfOrderHorizontal := model.Grid{}
			outOfOrderHorizontal.Carrier[0] = model.Coordinate{"A", 1}
			outOfOrderHorizontal.Carrier[1] = model.Coordinate{"A", 2}
			outOfOrderHorizontal.Carrier[2] = model.Coordinate{"A", 3}
			outOfOrderHorizontal.Carrier[3] = model.Coordinate{"A", 4}
			outOfOrderHorizontal.Carrier[4] = model.Coordinate{"A", 5}

			outOfOrderHorizontal.Battleship[0] = model.Coordinate{"A", 6}
			outOfOrderHorizontal.Battleship[1] = model.Coordinate{"A", 7}
			outOfOrderHorizontal.Battleship[2] = model.Coordinate{"A", 8}
			outOfOrderHorizontal.Battleship[3] = model.Coordinate{"A", 9}

			outOfOrderHorizontal.Cruiser[0] = model.Coordinate{"B", 1}
			outOfOrderHorizontal.Cruiser[1] = model.Coordinate{"B", 2}
			outOfOrderHorizontal.Cruiser[2] = model.Coordinate{"B", 3}

			outOfOrderHorizontal.Submarine[0] = model.Coordinate{"B", 4}
			outOfOrderHorizontal.Submarine[1] = model.Coordinate{"B", 5}
			outOfOrderHorizontal.Submarine[2] = model.Coordinate{"B", 6}

			outOfOrderHorizontal.Destroyer[0] = model.Coordinate{"B", 8}
			outOfOrderHorizontal.Destroyer[1] = model.Coordinate{"B", 7}

			gomega.Expect(allShipsInTouchingSpaces(outOfOrderHorizontal)).To(gomega.BeTrue())
		})

		ginkgo.It("flags as valid a model.Grid where all ships are in contiguous spaces, even if the model.Coordinates are not in order vertically", func() {
			outOfOrderVertical := model.Grid{}
			outOfOrderVertical.Carrier[0] = model.Coordinate{"A", 1}
			outOfOrderVertical.Carrier[1] = model.Coordinate{"A", 2}
			outOfOrderVertical.Carrier[2] = model.Coordinate{"A", 3}
			outOfOrderVertical.Carrier[3] = model.Coordinate{"A", 4}
			outOfOrderVertical.Carrier[4] = model.Coordinate{"A", 5}

			outOfOrderVertical.Battleship[0] = model.Coordinate{"A", 6}
			outOfOrderVertical.Battleship[1] = model.Coordinate{"A", 7}
			outOfOrderVertical.Battleship[2] = model.Coordinate{"A", 8}
			outOfOrderVertical.Battleship[3] = model.Coordinate{"A", 9}

			outOfOrderVertical.Cruiser[0] = model.Coordinate{"B", 1}
			outOfOrderVertical.Cruiser[1] = model.Coordinate{"B", 2}
			outOfOrderVertical.Cruiser[2] = model.Coordinate{"B", 3}

			outOfOrderVertical.Submarine[0] = model.Coordinate{"B", 4}
			outOfOrderVertical.Submarine[1] = model.Coordinate{"B", 5}
			outOfOrderVertical.Submarine[2] = model.Coordinate{"B", 6}

			outOfOrderVertical.Destroyer[0] = model.Coordinate{"C", 7}
			outOfOrderVertical.Destroyer[1] = model.Coordinate{"B", 7}

			gomega.Expect(allShipsInTouchingSpaces(outOfOrderVertical)).To(gomega.BeTrue())
		})

		ginkgo.It("flags as invalid a model.Grid where ships are in nonadjacent spaces", func() {
			invalid = model.Grid{}
			invalid.Carrier[0] = model.Coordinate{"A", 1}
			invalid.Carrier[1] = model.Coordinate{"A", 2}
			invalid.Carrier[2] = model.Coordinate{"A", 3}
			invalid.Carrier[3] = model.Coordinate{"A", 4}
			invalid.Carrier[4] = model.Coordinate{"A", 5}

			invalid.Battleship[0] = model.Coordinate{"A", 6}
			invalid.Battleship[1] = model.Coordinate{"A", 7}
			invalid.Battleship[2] = model.Coordinate{"A", 8}
			invalid.Battleship[3] = model.Coordinate{"A", 9}

			invalid.Cruiser[0] = model.Coordinate{"B", 1}
			invalid.Cruiser[1] = model.Coordinate{"B", 2}
			invalid.Cruiser[2] = model.Coordinate{"B", 3}

			invalid.Submarine[0] = model.Coordinate{"B", 4}
			invalid.Submarine[1] = model.Coordinate{"C", 5}
			invalid.Submarine[2] = model.Coordinate{"B", 6}

			invalid.Destroyer[0] = model.Coordinate{"B", 7}
			invalid.Destroyer[1] = model.Coordinate{"B", 9}

			gomega.Expect(allShipsInTouchingSpaces(invalid)).NotTo(gomega.BeTrue())
		})

		ginkgo.It("flags as invalid a model.Grid where a ship has one or more repeated spaces", func() {
			invalid = model.Grid{}
			invalid.Carrier[0] = model.Coordinate{"A", 1}
			invalid.Carrier[1] = model.Coordinate{"A", 2}
			invalid.Carrier[2] = model.Coordinate{"A", 3}
			invalid.Carrier[3] = model.Coordinate{"A", 4}
			invalid.Carrier[4] = model.Coordinate{"A", 5}

			invalid.Battleship[0] = model.Coordinate{"A", 6}
			invalid.Battleship[1] = model.Coordinate{"A", 7}
			invalid.Battleship[2] = model.Coordinate{"A", 8}
			invalid.Battleship[3] = model.Coordinate{"A", 9}

			invalid.Cruiser[0] = model.Coordinate{"B", 1}
			invalid.Cruiser[1] = model.Coordinate{"B", 2}
			invalid.Cruiser[2] = model.Coordinate{"B", 3}

			invalid.Submarine[0] = model.Coordinate{"B", 4}
			invalid.Submarine[1] = model.Coordinate{"C", 5}
			invalid.Submarine[2] = model.Coordinate{"B", 6}

			invalid.Destroyer[0] = model.Coordinate{"B", 7}
			invalid.Destroyer[1] = model.Coordinate{"B", 7}

			gomega.Expect(allShipsInTouchingSpaces(invalid)).NotTo(gomega.BeTrue())
		})
	})

	ginkgo.Describe("No ships overlap", func() {
		ginkgo.It("flags as valid a model.Grid where no ships overlap", func() {
			gomega.Expect(shipsDoNotOverlap(valid)).To(gomega.BeTrue())
		})

		ginkgo.It("flags as invalid a model.Grid where ships overlap", func() {
			invalid = model.Grid{}
			invalid.Carrier[0] = model.Coordinate{"A", 1}
			invalid.Carrier[1] = model.Coordinate{"A", 2}
			invalid.Carrier[2] = model.Coordinate{"A", 3}
			invalid.Carrier[3] = model.Coordinate{"A", 4}
			invalid.Carrier[4] = model.Coordinate{"A", 5}

			invalid.Battleship[0] = model.Coordinate{"A", 5}
			invalid.Battleship[1] = model.Coordinate{"A", 6}
			invalid.Battleship[2] = model.Coordinate{"A", 7}
			invalid.Battleship[3] = model.Coordinate{"A", 8}

			gomega.Expect(shipsDoNotOverlap(invalid)).NotTo(gomega.BeTrue())
		})
	})

	ginkgo.It("flags a model.Grid as valid if all criteria are met", func() {
		gomega.Expect(validateShipPlacement(valid)).To(gomega.BeTrue())
	})

	ginkgo.It("flags a model.Grid as invalid if a criterion is not met", func() {
		invalid = model.Grid{}
		invalid.Carrier[0] = model.Coordinate{"A", 1}
		invalid.Carrier[1] = model.Coordinate{"A", 2}
		invalid.Carrier[2] = model.Coordinate{"A", 3}
		invalid.Carrier[3] = model.Coordinate{"A", 4}
		invalid.Carrier[4] = model.Coordinate{"A", 5}

		invalid.Battleship[0] = model.Coordinate{"A", 6}
		invalid.Battleship[1] = model.Coordinate{"A", 7}
		invalid.Battleship[2] = model.Coordinate{"A", 8}
		invalid.Battleship[3] = model.Coordinate{"A", 9}

		invalid.Cruiser[0] = model.Coordinate{"B", 1}
		invalid.Cruiser[1] = model.Coordinate{"B", 2}
		invalid.Cruiser[2] = model.Coordinate{"B", 3}

		invalid.Submarine[0] = model.Coordinate{"B", 4}
		invalid.Submarine[1] = model.Coordinate{"C", 5}
		invalid.Submarine[2] = model.Coordinate{"B", 6}

		invalid.Destroyer[0] = model.Coordinate{"B", 7}
		invalid.Destroyer[1] = model.Coordinate{"B", 9}

		gomega.Expect(validateShipPlacement(invalid)).NotTo(gomega.BeTrue())
	})
})
