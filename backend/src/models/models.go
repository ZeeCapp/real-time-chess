package models

import "github.com/gorilla/websocket"

type Player struct {
	Name       string          `json:"name"`
	PlayerUUID string          `json:"playerUUID"`
	Socket     *websocket.Conn `json:"-"`
}

type Lobby struct {
	LobbyUUID string  `json:"lobbyUUID"`
	Player1   *Player `json:"player1"`
	Player2   *Player `json:"player2"`
}
