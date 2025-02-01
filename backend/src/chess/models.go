package chess

type PieceType string

const (
	Pawn   PieceType = "Pawn"
	Bishop PieceType = "Bishop"
	Knight PieceType = "Knight"
	Rook   PieceType = "Rook"
	Queen  PieceType = "Queen"
	King   PieceType = "King"
)

type Coordinates struct {
	X int8
	Y int8
}

type ChessPiece struct {
	Id         int8
	X          int8
	Y          int8
	Type       PieceType
	White      bool
	PlayerId   string
	Eliminated bool
}

type Player struct {
	Id       string
	Nickname string
	White    bool
}

/*
	the board is indexed like this:

	y (1-8)
	|
	|
	|
	|
	|
	--------------- x (a-h)
*/

type GameInstance struct {
	player1 Player
	player2 Player
	pieces  [32]ChessPiece
}
