package routers

import (
	"github.com/gin-gonic/gin"
)

func RegisterApiHandlers(group *gin.RouterGroup) {
	apiLobbyHandlersGroup := group.Group("/lobby")

	registerLobbyHandlers(apiLobbyHandlersGroup)
}
