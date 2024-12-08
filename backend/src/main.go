package main

import (
	"github.com/ZeeCapp/real-time-chess/src/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	apiHandlersGroup := r.Group("/api")

	routers.RegisterApiHandlers(apiHandlersGroup)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
