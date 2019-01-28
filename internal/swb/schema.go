package swb

import (
	"fmt"
	"strings"
)

// Coordinate one square on the grid
type Coordinate struct {
	Row    string `json:"row"`
	Column int    `json:"column"`
}

func (c Coordinate) String() string {
	return fmt.Sprintf("{%s,%d}", c.Row, c.Column)
}

// Grid placement of ships on the grid
type Grid struct {
	Carrier    [5]Coordinate `json:"carrier"`
	Battleship [4]Coordinate `json:"battleship"`
	Cruiser    [3]Coordinate `json:"cruiser"`
	Submarine  [3]Coordinate `json:"submarine"`
	Destroyer  [2]Coordinate `json:"destroyer"`
}

func (g Grid) String() string {
	var gridString string
	gridString = strings.Join(shipToString(g.Carrier[:]), ",") + "\n"
	gridString = gridString + strings.Join(shipToString(g.Battleship[:]), ",") + "\n"
	gridString = gridString + strings.Join(shipToString(g.Cruiser[:]), ",") + "\n"
	gridString = gridString + strings.Join(shipToString(g.Submarine[:]), ",") + "\n"
	gridString = gridString + strings.Join(shipToString(g.Destroyer[:]), ",")
	return gridString
}

func shipToString(ship []Coordinate) []string {
	var shipStrings []string
	for _, coord := range ship {
		shipStrings = append(shipStrings, coord.String())
	}
	return shipStrings
}

// Player combination of ships and shots
type Player struct {
	Ships Grid          `json:"ships"`
	Shots *[]Coordinate `json:"shots"`
}

// Game full state of a Starish Wars Battleship game
type Game struct {
	Player1 Player     `json:"player_1"`
	Player2 Player     `json:"player_2"`
	Status  GameStatus `json:"status"`
}

type GameStatus string

const (
	AWAITING_SHIPS  GameStatus = "AWAITING_SHIPS"
	PLAYER_1_ACTIVE GameStatus = "PLAYER_1_ACTIVE"
	PLAYER_2_ACTIVE GameStatus = "PLAYER_2_ACTIVE"
	GAME_OVER       GameStatus = "GAME_OVER"
)
