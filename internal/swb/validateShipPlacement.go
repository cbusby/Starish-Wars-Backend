package swb

func validateShipPlacement(grid Grid) bool {
	return allShipsOnBoard(grid) &&
		allShipsHorizontalOrVertical(grid) &&
		allShipsInTouchingSpaces(grid) &&
		shipsDoNotOverlap(grid)
}

func allShipsOnBoard(grid Grid) bool {
	var retVal = true
	for i := 0; i < len(grid.Carrier); i++ {
		if grid.Carrier[i] == (Coordinate{}) {
			retVal = false
		}
	}
	for i := 0; i < len(grid.Battleship); i++ {
		if grid.Battleship[i] == (Coordinate{}) {
			retVal = false
		}
	}
	for i := 0; i < len(grid.Cruiser); i++ {
		if grid.Cruiser[i] == (Coordinate{}) {
			retVal = false
		}
	}
	for i := 0; i < len(grid.Submarine); i++ {
		if grid.Submarine[i] == (Coordinate{}) {
			retVal = false
		}
	}
	for i := 0; i < len(grid.Destroyer); i++ {
		if grid.Destroyer[i] == (Coordinate{}) {
			retVal = false
		}
	}
	return retVal
}

func allShipsHorizontalOrVertical(grid Grid) bool {
	return false
}

func allShipsInTouchingSpaces(grid Grid) bool {
	return false
}

func shipsDoNotOverlap(grid Grid) bool {
	return false
}
