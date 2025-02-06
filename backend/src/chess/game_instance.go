package chess

import (
	"errors"

	"github.com/ZeeCapp/real-time-chess/src/helpers"
)

func (gameInstance *GameInstance) initDefaultState() {
	var whitePlayerId, blackPlayerId string

	if gameInstance.player1.White {
		whitePlayerId = gameInstance.player1.Id
		blackPlayerId = gameInstance.player2.Id
	} else {
		whitePlayerId = gameInstance.player2.Id
		blackPlayerId = gameInstance.player1.Id
	}

	initSpecialPieces(&gameInstance.pieces, 24, blackPlayerId, 7, false)
	initPawns(&gameInstance.pieces, 16, blackPlayerId, 6, false)

	initPawns(&gameInstance.pieces, 8, whitePlayerId, 1, true)
	initSpecialPieces(&gameInstance.pieces, 0, whitePlayerId, 0, true)

	gameInstance.enPassantPawns = make([]ChessPiece, 0, 16)
}

func (gameInstance *GameInstance) GetPieces() [32]ChessPiece {
	return gameInstance.pieces
}

func (gameInstance *GameInstance) findPieceAtPosition(x int8, y int8) *ChessPiece {
	for _, piece := range gameInstance.pieces {
		if piece.X == x && piece.Y == y {
			return &piece
		}
	}

	return nil
}

func (gameInstance *GameInstance) Move(pieceId int8, endX int8, endY int8) error {
	var pieceToMove *ChessPiece

	for _, piece := range gameInstance.pieces {
		if piece.Id == pieceId {
			pieceToMove = &piece
		}
	}

	if pieceToMove == nil {
		return (errors.New("invalid piece id"))
	}

	if !isInsideBoardBoundaries(endX, endY) {
		return errors.New("end position outsode of board dimensions")
	}

	var moveErr error

	switch pieceToMove.Type {
	case Pawn:
		moveErr = movePawn(gameInstance, pieceToMove, endX, endY)
	case Bishop:
		moveErr = moveBishop(gameInstance, pieceToMove, endX, endY)
	case Knight:
		moveErr = moveKnight(gameInstance, pieceToMove, endX, endY)
	case Rook:
		moveErr = moveRook(gameInstance, pieceToMove, endX, endY)
	case Queen:
		moveErr = moveQueen(gameInstance, pieceToMove, endX, endY)
	case King:
		moveErr = moveKing(gameInstance, pieceToMove, endX, endY)
	}

	return moveErr
}

// TODO: en passant moves for pawns
func movePawn(gameInstance *GameInstance, piece *ChessPiece, endX int8, endY int8) error {
	if endX == piece.X && endY == piece.Y {
		return nil
	}

	var validTargetCoords []*Coordinates = make([]*Coordinates, 6)

	// set all valid coords to -128 instead of 0 to avoid valid moves at x:0 & y:0 because
	// they are inside the playing board
	for _, validCoords := range validTargetCoords {
		validCoords.X = -128
		validCoords.Y = -128
	}

	var diagonalLeftPiece *ChessPiece
	var diagonalRightPiece *ChessPiece
	currentX := piece.X
	currentY := piece.Y

	appendPawnForwardMoves(gameInstance, piece, validTargetCoords)

	if piece.White {
		diagonalLeftPiece = gameInstance.findPieceAtPosition(currentX-1, currentY+1)
		diagonalRightPiece = gameInstance.findPieceAtPosition(currentX+1, currentY+1)
	} else {
		diagonalLeftPiece = gameInstance.findPieceAtPosition(currentX-1, currentY-1)
		diagonalRightPiece = gameInstance.findPieceAtPosition(currentX+1, currentY-1)
	}

	if diagonalLeftPiece != nil && !isFirendly(diagonalLeftPiece, piece) {
		if piece.White {
			validTargetCoords = append(validTargetCoords, &Coordinates{X: currentX - 1, Y: currentY + 1})
		} else {
			validTargetCoords = append(validTargetCoords, &Coordinates{X: currentX - 1, Y: currentY - 1})
		}
	}

	if diagonalRightPiece != nil && !isFirendly(diagonalRightPiece, piece) {
		if piece.White {
			validTargetCoords = append(validTargetCoords, &Coordinates{X: currentX + 1, Y: currentY + 1})
		} else {
			validTargetCoords = append(validTargetCoords, &Coordinates{X: currentX + 1, Y: currentY - 1})
		}
	}

	isValidMove := false

	for _, validCoords := range validTargetCoords {
		if validCoords.X == endX && validCoords.Y == endY {
			isValidMove = true
			break
		}
	}

	if !isValidMove {
		return errors.New("invalid pawn move")
	}

	gameInstance.findPieceAtPosition(endX, endY).Eliminated = true

	piece.X = endX
	piece.Y = endY

	if endY-currentY == 2 || endY-currentY == -2 {
		gameInstance.enPassantPawns = append(gameInstance.enPassantPawns, *piece)
	}

	return nil
}

func moveBishop(gameInstance *GameInstance, piece *ChessPiece, endX int8, endY int8) error {
	return nil
}

func moveKnight(gameInstance *GameInstance, piece *ChessPiece, endX int8, endY int8) error {
	return nil
}

func moveRook(gameInstance *GameInstance, piece *ChessPiece, endX int8, endY int8) error {
	return nil
}

func moveQueen(gameInstance *GameInstance, piece *ChessPiece, endX int8, endY int8) error {
	return nil
}

/*
func handleEnPassant(gameInstance *GameInstance, piece *ChessPiece) {
	for _, enPassantPawn := range gameInstance.enPassantPawns {

	}
}
*/

func clearEnPassantPawns(gameInstance *GameInstance, piece *ChessPiece) {
	leftPieces := make([]ChessPiece, 0, 12)

	for _, enPassantPiece := range gameInstance.enPassantPawns {
		if piece.White != enPassantPiece.White {
			leftPieces = append(leftPieces, enPassantPiece)
		}
	}

	gameInstance.enPassantPawns = leftPieces
}

func moveKing(gameInstance *GameInstance, piece *ChessPiece, endX int8, endY int8) error {
	if endX == piece.X && endY == piece.Y {
		return nil
	}

	if !isInsideBoardBoundaries(endX, endY) {
		return errors.New("move outside of board boundaries")
	}

	var validYMoves = []int8{1, 0, -1}
	var validXMoves = []int8{1, 0, -1}

	currentX := piece.X
	currentY := piece.Y

	if !helpers.In(validYMoves, currentY-endY) {
		return errors.New("forbidden king move")
	}

	if !helpers.In(validXMoves, currentX-endX) {
		return errors.New("forbidden king move")
	}

	targetPosPiece := gameInstance.findPieceAtPosition(endX, endY)

	if (targetPosPiece.White && piece.White) || (!targetPosPiece.White && !piece.White) {
		return errors.New("end position allready ocupied by friendly piece")
	}

	targetPosPiece.Eliminated = true

	piece.X = endX
	piece.Y = endY

	return nil
}

func appendPawnForwardMoves(game *GameInstance, pawn *ChessPiece, coordinates []*Coordinates) {
	for i := 0; i < 2; i++ {
		moveCoords := getRelativeMove(pawn, int8(i), 0, 0, 0)
		pieceAtCoords := game.findPieceAtPosition(moveCoords.X, moveCoords.Y)

		if pieceAtCoords != nil {
			return
		}

		coordinates = append(coordinates, &moveCoords)
	}
}

func getRelativeMove(piece *ChessPiece, forward int8, back int8, left int8, right int8) Coordinates {
	if piece.White {
		return Coordinates{
			X: piece.X + left - right,
			Y: piece.Y + forward - back,
		}
	} else {
		return Coordinates{
			X: piece.X - left + right,
			Y: piece.Y - forward + back,
		}
	}
}

func isFirendly(piece1 *ChessPiece, piece2 *ChessPiece) bool {
	if (piece1.White && !piece2.White) || (!piece1.White && piece2.White) {
		return true
	}
	return false
}

func isInsideBoardBoundaries(x int8, y int8) bool {
	if x < 0 || x > 7 {
		return false
	}

	if y < 0 || y > 7 {
		return false
	}

	return true
}

func initSpecialPieces(pieces *[32]ChessPiece, startId int, playerId string, yPosition int8, white bool) {
	var currentX int8

	for i := startId; i < startId+8; i++ {
		piece := &pieces[i]

		piece.Id = int8(i + 1)
		piece.X = currentX
		piece.Y = yPosition
		piece.White = white
		piece.PlayerId = playerId
		piece.Eliminated = false

		switch i {
		case 0, 7:
			piece.Type = Rook
		case 1, 6:
			piece.Type = Knight
		case 2, 5:
			piece.Type = Bishop
		case 3:
			piece.Type = Queen
		case 4:
			piece.Type = King
		}

		currentX++
	}
}

func initPawns(pieces *[32]ChessPiece, startId int, playerId string, yPosition int8, white bool) {
	var currentX int8

	for i := startId; i < startId+8; i++ {
		piece := &pieces[i]

		piece.Id = int8(i + 1)
		piece.X = currentX
		piece.Y = yPosition
		piece.Type = Pawn
		piece.White = white
		piece.PlayerId = playerId
		piece.Eliminated = false

		currentX++
	}
}
