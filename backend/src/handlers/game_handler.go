package handlers

import (
	"github.com/ZeeCapp/real-time-chess/src/loggers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleSocketUpgrade(contex *gin.Context) {
	_, err := upgrader.Upgrade(contex.Writer, contex.Request, nil)

	if err != nil {
		loggers.GetLoggerInstance().LogError(err.Error())
	}

	// TODO: Implement game logic
}
