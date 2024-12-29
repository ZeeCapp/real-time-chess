package handlers

import (
	"github.com/ZeeCapp/real-time-chess/src/loggers"
	"github.com/ZeeCapp/real-time-chess/src/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type socketUpgradeUri struct {
	LobbyUUID  string `uri:"lobbyUUID" binding:"required,uuid"`
	PlayerUUID string `uri:"playerUUID" binding:"required,uuid"`
}

func HandleSocketUpgrade(contex *gin.Context) {

	var requestUri socketUpgradeUri

	bindingErr := contex.ShouldBindUri(&requestUri)

	if bindingErr != nil {
		loggers.GetLoggerInstance().LogError(bindingErr.Error())
		contex.Status(400)
		return
	}

	var lobby *models.Lobby
	var player *models.Player

	for _, l := range lobbies {
		if l.LobbyUUID == requestUri.LobbyUUID {
			lobby = &l
		}
	}

	if lobby == nil {
		contex.Status(404)
		return
	}

	if lobby.Player1.PlayerUUID == requestUri.PlayerUUID {
		player = lobby.Player1
	} else if lobby.Player2.PlayerUUID == requestUri.PlayerUUID {
		player = lobby.Player2
	} else {
		contex.Status(404)
		return
	}

	socket, err := upgrader.Upgrade(contex.Writer, contex.Request, nil)

	if err != nil {
		loggers.GetLoggerInstance().LogError(err.Error())
	}

	player.Socket = socket
}
