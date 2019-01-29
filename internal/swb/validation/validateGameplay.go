package validation

import (
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
	firstCopy := model.SortAndCopyShip(firstShip)
	secondCopy := model.SortAndCopyShip(secondShip)
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

// AllShipsSunk check if all of a player's ships have been sunk
func AllShipsSunk(inactivePlayer model.Player, activePlayerShots []model.Coordinate) bool {
	ships := inactivePlayer.Ships
	if !shipIsSunk(ships.Carrier[:], activePlayerShots) {
		return false
	}
	if !shipIsSunk(ships.Battleship[:], activePlayerShots) {
		return false
	}
	if !shipIsSunk(ships.Cruiser[:], activePlayerShots) {
		return false
	}
	if !shipIsSunk(ships.Submarine[:], activePlayerShots) {
		return false
	}
	if !shipIsSunk(ships.Destroyer[:], activePlayerShots) {
		return false
	}
	return true
}

func shipIsSunk(ship []model.Coordinate, shots []model.Coordinate) bool {
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
