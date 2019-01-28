package model

// Coordinate one square on the grid
type Coordinate struct {
	Row    string `json:"row"`
	Column int    `json:"column"`
}

// Grid placement of ships on the grid
type Grid struct {
	Carrier    [5]Coordinate `json:"carrier"`
	Battleship [4]Coordinate `json:"battleship"`
	Cruiser    [3]Coordinate `json:"cruiser"`
	Submarine  [3]Coordinate `json:"submarine"`
	Destroyer  [2]Coordinate `json:"destroyer"`
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
