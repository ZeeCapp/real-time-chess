package routers

import (
	"github.com/ZeeCapp/real-time-chess/src/handlers"
	"github.com/gin-gonic/gin"
)

func registerLobbyHandlers(group *gin.RouterGroup) {
	group.GET("/create", handlers.HandleLobbyCreation)

	group.GET("/join", handlers.HandleLobbyJoin)
}
