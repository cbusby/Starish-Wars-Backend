package validation

import "github.com/cbusby/Starish-Wars-Backend/internal/swb/model"

// ShipsSame validate that a player's ships are the same from turn to turn
func ShipsSame(oldPlayer model.Player, newPlayer model.Player) bool {
	return false
}

// ShotsSame validate that the inactive player's shot list are the same from turn to turn
func ShotsSame(oldPlayer model.Player, newPlayer model.Player) bool {
	return false
}

// OneNewShot validate that they active player's shot list is the same except for one new shot
func OneNewShot(oldPlayer model.Player, newPlayer model.Player) bool {
	return false
}

// AllShipsHit check if all of a player's ships have been sunk
func AllShipsHit(inactivePlayer model.Player, activePlayerShots *[]model.Coordinate) bool {
	return false
}
