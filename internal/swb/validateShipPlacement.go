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
	var retVal = true
	retVal = retVal && (shipIsHorizontal(grid.Carrier[:]) || shipIsVertical(grid.Carrier[:]))
	retVal = retVal && (shipIsHorizontal(grid.Battleship[:]) || shipIsVertical(grid.Battleship[:]))
	retVal = retVal && (shipIsHorizontal(grid.Cruiser[:]) || shipIsVertical(grid.Cruiser[:]))
	retVal = retVal && (shipIsHorizontal(grid.Submarine[:]) || shipIsVertical(grid.Submarine[:]))
	retVal = retVal && (shipIsHorizontal(grid.Destroyer[:]) || shipIsVertical(grid.Destroyer[:]))
	return retVal
}

func shipIsHorizontal(ship []Coordinate) bool {
	for i := 1; i < len(ship); i++ {
		if ship[i].Column != ship[0].Column {
			return false
		}
	}
	return true
}

func shipIsVertical(ship []Coordinate) bool {
	for i := 1; i < len(ship); i++ {
		if ship[i].Row != ship[0].Row {
			return false
		}
	}
	return true
}

func allShipsInTouchingSpaces(grid Grid) bool {
	return false
}

func shipsDoNotOverlap(grid Grid) bool {
	return false
}
