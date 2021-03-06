package validation

import (
	"github.com/cbusby/Starish-Wars-Backend/internal/swb/model"
)

// ValidateShipPlacement perform all validations for correct ship placement
func ValidateShipPlacement(grid model.Grid) bool {
	return AllShipsPresent(grid) &&
		allShipsOnGrid(grid) &&
		allShipsHorizontalOrVertical(grid) &&
		allShipsInTouchingSpaces(grid) &&
		shipsDoNotOverlap(grid)
}

// AllShipsPresent validate that all ships are placed on grid
func AllShipsPresent(grid model.Grid) bool {
	for i := 0; i < len(grid.Carrier); i++ {
		if grid.Carrier[i] == (model.Coordinate{}) {
			return false
		}
	}
	for i := 0; i < len(grid.Battleship); i++ {
		if grid.Battleship[i] == (model.Coordinate{}) {
			return false
		}
	}
	for i := 0; i < len(grid.Cruiser); i++ {
		if grid.Cruiser[i] == (model.Coordinate{}) {
			return false
		}
	}
	for i := 0; i < len(grid.Submarine); i++ {
		if grid.Submarine[i] == (model.Coordinate{}) {
			return false
		}
	}
	for i := 0; i < len(grid.Destroyer); i++ {
		if grid.Destroyer[i] == (model.Coordinate{}) {
			return false
		}
	}
	return true
}

func allShipsOnGrid(grid model.Grid) bool {
	if !allCoordsOnGrid(grid.Carrier[:]) {
		return false
	}
	if !allCoordsOnGrid(grid.Battleship[:]) {
		return false
	}
	if !allCoordsOnGrid(grid.Cruiser[:]) {
		return false
	}
	if !allCoordsOnGrid(grid.Submarine[:]) {
		return false
	}
	if !allCoordsOnGrid(grid.Destroyer[:]) {
		return false
	}
	return true
}

func allCoordsOnGrid(ship []model.Coordinate) bool {
	for i := 0; i < len(ship); i++ {
		if ship[i].Row[0] < 'A' || ship[i].Row[0] > 'J' || ship[i].Column < 1 || ship[i].Column > 10 {
			return false
		}
	}
	return true
}

func allShipsHorizontalOrVertical(grid model.Grid) bool {
	if !(shipIsHorizontal(grid.Carrier[:]) || shipIsVertical(grid.Carrier[:])) {
		return false
	}
	if !(shipIsHorizontal(grid.Battleship[:]) || shipIsVertical(grid.Battleship[:])) {
		return false
	}
	if !(shipIsHorizontal(grid.Cruiser[:]) || shipIsVertical(grid.Cruiser[:])) {
		return false
	}
	if !(shipIsHorizontal(grid.Submarine[:]) || shipIsVertical(grid.Submarine[:])) {
		return false
	}
	if !(shipIsHorizontal(grid.Destroyer[:]) || shipIsVertical(grid.Destroyer[:])) {
		return false
	}
	return true
}

func shipIsHorizontal(ship []model.Coordinate) bool {
	for i := 1; i < len(ship); i++ {
		if ship[i].Row != ship[0].Row {
			return false
		}
	}
	return true
}

func shipIsVertical(ship []model.Coordinate) bool {
	for i := 1; i < len(ship); i++ {
		if ship[i].Column != ship[0].Column {
			return false
		}
	}
	return true
}

func allShipsInTouchingSpaces(grid model.Grid) bool {
	if !shipIsContiguous(grid.Carrier[:]) {
		return false
	}
	if !shipIsContiguous(grid.Battleship[:]) {
		return false
	}
	if !shipIsContiguous(grid.Cruiser[:]) {
		return false
	}
	if !shipIsContiguous(grid.Submarine[:]) {
		return false
	}
	if !shipIsContiguous(grid.Destroyer[:]) {
		return false
	}
	return true
}

func shipIsContiguous(ship []model.Coordinate) bool {
	copy := model.SortAndCopyShip(ship)
	if shipIsHorizontal(copy) {
		for i := 1; i < len(copy); i++ {
			if copy[i].Column-copy[i-1].Column != 1 {
				return false
			}
		}
	} else {
		for i := 1; i < len(copy); i++ {
			if copy[i].Row[0]-copy[i-1].Row[0] != 1 {
				return false
			}
		}
	}
	return true
}

func shipsDoNotOverlap(grid model.Grid) bool {
	if shipsOverlap(grid.Carrier[:], grid.Battleship[:]) {
		return false
	}
	if shipsOverlap(grid.Carrier[:], grid.Cruiser[:]) {
		return false
	}
	if shipsOverlap(grid.Carrier[:], grid.Submarine[:]) {
		return false
	}
	if shipsOverlap(grid.Carrier[:], grid.Destroyer[:]) {
		return false
	}
	if shipsOverlap(grid.Battleship[:], grid.Cruiser[:]) {
		return false
	}
	if shipsOverlap(grid.Battleship[:], grid.Submarine[:]) {
		return false
	}
	if shipsOverlap(grid.Battleship[:], grid.Destroyer[:]) {
		return false
	}
	if shipsOverlap(grid.Cruiser[:], grid.Submarine[:]) {
		return false
	}
	if shipsOverlap(grid.Cruiser[:], grid.Destroyer[:]) {
		return false
	}
	if shipsOverlap(grid.Submarine[:], grid.Destroyer[:]) {
		return false
	}
	return true
}

func shipsOverlap(ship1 []model.Coordinate, ship2 []model.Coordinate) bool {
	for i := 0; i < len(ship2); i++ {
		for j := 0; j < len(ship1); j++ {
			if ship1[j] == ship2[i] {
				return true
			}
		}
	}
	return false
}
