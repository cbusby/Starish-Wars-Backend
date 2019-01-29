package validation

import (
	"sort"

	"github.com/cbusby/Starish-Wars-Backend/internal/swb/model"
)

// ShipsSame validate that a player's ships are the same from turn to turn
func ShipsSame(oldPlayer model.Player, newPlayer model.Player) bool {
	oldGrid := oldPlayer.Ships
	newGrid := newPlayer.Ships
	if !shipsAreSame(oldGrid.Carrier[:], newGrid.Carrier[:]) {
		return false
	}
	if !shipsAreSame(oldGrid.Battleship[:], newGrid.Battleship[:]) {
		return false
	}
	if !shipsAreSame(oldGrid.Cruiser[:], newGrid.Cruiser[:]) {
		return false
	}
	if !shipsAreSame(oldGrid.Submarine[:], newGrid.Submarine[:]) {
		return false
	}
	if !shipsAreSame(oldGrid.Destroyer[:], newGrid.Destroyer[:]) {
		return false
	}
	return true
}

func shipsAreSame(firstShip []model.Coordinate, secondShip []model.Coordinate) bool {
	if len(firstShip) != len(secondShip) {
		return false
	}
	firstCopy := append([]model.Coordinate(nil), firstShip...)
	secondCopy := append([]model.Coordinate(nil), secondShip...)
	sort.Slice(firstCopy, func(i, j int) bool {
		if firstCopy[i].Row == firstCopy[j].Row {
			return firstCopy[i].Column < firstCopy[j].Column
		}
		return firstCopy[i].Row < firstCopy[j].Row
	})
	sort.Slice(secondCopy, func(i, j int) bool {
		if secondCopy[i].Row == secondCopy[j].Row {
			return secondCopy[i].Column < secondCopy[j].Column
		}
		return secondCopy[i].Row < secondCopy[j].Row
	})
	for i := 0; i < len(firstCopy); i++ {
		if firstCopy[i] != secondCopy[i] {
			return false
		}
	}
	return true
}

// ShotsSame validate that the inactive player's shot list are the same from turn to turn
func ShotsSame(oldPlayer model.Player, newPlayer model.Player) bool {
	oldShots := oldPlayer.Shots
	newShots := newPlayer.Shots
	if len(oldShots) != len(newShots) {
		return false
	}
	for i := 0; i < len(oldShots); i++ {
		if oldShots[i] != newShots[i] {
			return false
		}
	}
	return true
}

// OneNewShot validate that the active player's shot list is the same except for one new shot
func OneNewShot(oldPlayer model.Player, newPlayer model.Player) bool {
	oldShots := oldPlayer.Shots
	newShots := newPlayer.Shots
	if len(oldShots)+1 != len(newShots) {
		return false
	}
	lastNewShot := newShots[len(newShots)-1]
	if lastNewShot.Row[0] < 'A' || lastNewShot.Row[0] > 'J' || lastNewShot.Column < 1 || lastNewShot.Column > 10 {
		return false
	}
	for i := 0; i < len(oldShots); i++ {
		if oldShots[i] != newShots[i] {
			return false
		}
		if oldShots[i] == lastNewShot {
			return false
		}
	}
	return true
}

// AllShipsHit check if all of a player's ships have been sunk
func AllShipsHit(inactivePlayer model.Player, activePlayerShots []model.Coordinate) bool {
	ships := inactivePlayer.Ships
	if !shipHasBeenSunk(ships.Carrier[:], activePlayerShots) {
		return false
	}
	if !shipHasBeenSunk(ships.Battleship[:], activePlayerShots) {
		return false
	}
	if !shipHasBeenSunk(ships.Cruiser[:], activePlayerShots) {
		return false
	}
	if !shipHasBeenSunk(ships.Submarine[:], activePlayerShots) {
		return false
	}
	if !shipHasBeenSunk(ships.Destroyer[:], activePlayerShots) {
		return false
	}
	return true
}

func shipHasBeenSunk(ship []model.Coordinate, shots []model.Coordinate) bool {
	spacesHit := 0
	for i := 0; i < len(shots); i++ {
		for j := 0; j < len(ship); j++ {
			if shots[i] == ship[j] {
				spacesHit++
				break
			}
		}
	}
	return spacesHit == len(ship)
}
