package model

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Schema", func() {
	ginkgo.Describe("Coordinate slice", func() {

		var (
			orig []Coordinate
			copy []Coordinate
		)

		ginkgo.BeforeEach(func() {
			orig := []Coordinate{
				Coordinate{"B", 1},
				Coordinate{"A", 3},
				Coordinate{"A", 2},
			}
			copy = SortAndCopyShip(orig)
		})

		ginkgo.It("copies an array slice of Coordinates, does not sort in place", func() {
			gomega.Expect(&orig).NotTo(gomega.Equal(&copy))
		})

		ginkgo.It("correctly sorts an array slice of Coordinates", func() {
			expected := []Coordinate{
				Coordinate{"A", 2},
				Coordinate{"A", 3},
				Coordinate{"B", 1},
			}
			gomega.Expect(copy).To(gomega.Equal(expected))
		})
	})
})
