package routers

import (
	"github.com/ZeeCapp/real-time-chess/src/handlers"
	"github.com/gin-gonic/gin"
)

func registerLobbyHandlers(group *gin.RouterGroup) {
	group.POST("/create", handlers.HandleLobbyCreation)

	group.POST("/:lobbyUUID/join", handlers.HandleLobbyJoin)

	group.GET("/:lobbyUUID/game/:playerUUID/join", handlers.HandleSocketUpgrade)
}
