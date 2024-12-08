package models

import "github.com/gorilla/websocket"

type Player struct {
	Name   string
	Guid   string
	Socket websocket.Conn
}

type Lobby struct {
	Guid    string
	Player1 *Player
	Player2 *Player
}
