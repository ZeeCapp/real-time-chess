package handlers

import (
	"github.com/ZeeCapp/real-time-chess/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var lobbies []models.Lobby = make([]models.Lobby, 0)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleLobbyCreation(context *gin.Context) {
	lobby := models.Lobby{Guid: uuid.NewString()}

	socket, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		context.Status(400)
		return
	}

	lobby.Player1 = &models.Player{
		Name: "Test",
		Guid: uuid.NewString(),
	}

	lobbies = append(lobbies, lobby)

	socket.WriteMessage(0, []byte("Connected"))
}

func HandleLobbyJoin(context *gin.Context) {

}
