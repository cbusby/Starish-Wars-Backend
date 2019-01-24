package swb

import (
	"sort"
)

func validateShipPlacement(grid Grid) bool {
	return allShipsPresent(grid) &&
		allShipsOnGrid(grid) &&
		allShipsHorizontalOrVertical(grid) &&
		allShipsInTouchingSpaces(grid) &&
		shipsDoNotOverlap(grid)
}

func allShipsPresent(grid Grid) bool {
	for i := 0; i < len(grid.Carrier); i++ {
		if grid.Carrier[i] == (Coordinate{}) {
			return false
		}
	}
	for i := 0; i < len(grid.Battleship); i++ {
		if grid.Battleship[i] == (Coordinate{}) {
			return false
		}
	}
	for i := 0; i < len(grid.Cruiser); i++ {
		if grid.Cruiser[i] == (Coordinate{}) {
			return false
		}
	}
	for i := 0; i < len(grid.Submarine); i++ {
		if grid.Submarine[i] == (Coordinate{}) {
			return false
		}
	}
	for i := 0; i < len(grid.Destroyer); i++ {
		if grid.Destroyer[i] == (Coordinate{}) {
			return false
		}
	}
	return true
}

func allShipsOnGrid(grid Grid) bool {
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

func allCoordsOnGrid(ship []Coordinate) bool {
	for i := 0; i < len(ship); i++ {
		if ship[i].Row < 'A' || ship[i].Row > 'J' || ship[i].Column < 1 || ship[i].Column > 10 {
			return false
		}
	}
	return true
}

func allShipsHorizontalOrVertical(grid Grid) bool {
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

func shipIsHorizontal(ship []Coordinate) bool {
	for i := 1; i < len(ship); i++ {
		if ship[i].Row != ship[0].Row {
			return false
		}
	}
	return true
}

func shipIsVertical(ship []Coordinate) bool {
	for i := 1; i < len(ship); i++ {
		if ship[i].Column != ship[0].Column {
			return false
		}
	}
	return true
}

func allShipsInTouchingSpaces(grid Grid) bool {
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

func shipIsContiguous(ship []Coordinate) bool {
	sort.Slice(ship, func(i, j int) bool {
		if ship[i].Row == ship[j].Row {
			return ship[i].Column < ship[j].Column
		} else {
			return ship[i].Row < ship[j].Row
		}
	})
	if shipIsHorizontal(ship) {
		for i := 1; i < len(ship); i++ {
			if ship[i].Column-ship[i-1].Column > 1 {
				return false
			}
		}
	} else {
		for i := 1; i < len(ship); i++ {
			if ship[i].Row-ship[i-1].Row > 1 {
				return false
			}
		}
	}
	return true
}

func shipsDoNotOverlap(grid Grid) bool {
	return false
}
