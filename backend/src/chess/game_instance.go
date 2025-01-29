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

	initSpecialPieces(&gameInstance.pieces, 0, whitePlayerId, 0, true)
	initPawns(&gameInstance.pieces, 8, whitePlayerId, 1, true)

	initPawns(&gameInstance.pieces, 16, blackPlayerId, 6, false)
	initSpecialPieces(&gameInstance.pieces, 24, blackPlayerId, 7, false)
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
	case Knight:
	case Rook:
	case Queen:
	case King:
	}

	return moveErr
}

func movePawn(gameInstance *GameInstance, piece *ChessPiece, endX int8, endY int8) error {
	var validYMove int8
	var validXMoves = []int8{1, 0, -1}

	if piece.White {
		validYMove = 1
	} else {
		validYMove = -1
	}

	currentX := piece.X
	currentY := piece.Y

	if currentY-endY != validYMove {
		return errors.New("forbidden pawn move")
	}

	if !helpers.In(validXMoves, currentX-endX) {
		return errors.New("forbidden pawn move")
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
