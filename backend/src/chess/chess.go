package chess

import (
	"errors"
)

func NewGame(player1 Player, player2 Player) (*GameInstance, error) {

	if player1.White && player2.White {
		return nil, errors.New("one player needs to be playing black")
	}

	newGame := GameInstance{
		player1: player1,
		player2: player2,
	}

	newGame.initDefaultState()

	return &newGame, nil
}
