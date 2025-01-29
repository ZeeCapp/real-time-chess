package main

import (
	"encoding/json"
	"fmt"

	"github.com/ZeeCapp/real-time-chess/src/chess"
	"github.com/ZeeCapp/real-time-chess/src/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Test struct {
	Message string
}

func main() {
	r := gin.Default()

	game, err := chess.NewGame(
		chess.Player{
			Id:    "Player1",
			White: true,
		},
		chess.Player{
			Id: "Player2",
		})

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%s\n\n", prettyPrint(game.GetPieces()))

	apiHandlersGroup := r.Group("/api")

	routers.RegisterApiHandlers(apiHandlersGroup)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func prettyPrint(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "\t")

	if err != nil {
		panic(err)
	}

	return string(s)
}
