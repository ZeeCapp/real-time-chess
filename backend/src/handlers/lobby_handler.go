package handlers

import (
	"fmt"

	"github.com/ZeeCapp/real-time-chess/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userRequestBody struct {
	Username string `json:"username"`
}

type lobbyCreatedResponse struct {
	Lobby           models.Lobby `json:"lobby"`
	GameJoinEndoint string       `json:"gameJoinEndpoint"`
	LobbyJoinURL    string       `json:"lobbyJoinURL"`
}

type joinLobbyRequestParams struct {
	LobbyUUID string `uri:"lobbyUUID" binding:"required,uuid"`
}

var lobbies []models.Lobby = make([]models.Lobby, 0)

func HandleLobbyCreation(context *gin.Context) {

	lobby := models.Lobby{LobbyUUID: uuid.NewString()}

	var requestPlayer userRequestBody

	bodyErr := context.ShouldBindBodyWithJSON(&requestPlayer)

	if bodyErr != nil {
		context.Status(400)
		return
	}

	lobby.Player1 = &models.Player{
		Name:       requestPlayer.Username,
		PlayerUUID: uuid.NewString(),
	}

	lobbies = append(lobbies, lobby)

	var responseBody lobbyCreatedResponse = lobbyCreatedResponse{
		Lobby:           lobby,
		GameJoinEndoint: fmt.Sprintf("/api/lobby/%s/game/join", lobby.LobbyUUID),
		LobbyJoinURL:    fmt.Sprintf("/invite/%s", lobby.LobbyUUID),
	}

	context.JSON(200, responseBody)
}

func HandleLobbyJoin(context *gin.Context) {

	var lobbyJoinRequest joinLobbyRequestParams
	var requestPlayer userRequestBody

	uriErr := context.ShouldBindUri(&lobbyJoinRequest)
	bodyErr := context.ShouldBindBodyWithJSON(&requestPlayer)

	if uriErr != nil || bodyErr != nil {
		context.Status(400)
		return
	}

	for _, lobby := range lobbies {
		if lobby.LobbyUUID == lobbyJoinRequest.LobbyUUID {
			lobby.Player2 = &models.Player{
				Name:       requestPlayer.Username,
				PlayerUUID: uuid.NewString(),
			}
			context.JSON(200, lobby)
			return
		}
	}

	context.Status(404)
}
