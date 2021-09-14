package app

import (
	"github.com/Stezok/game-tools/internal/redactor/character"
	"github.com/Stezok/game-tools/internal/redactor/item"
	"github.com/gin-gonic/gin"
)

type Logger interface {
	Print(...interface{})
}

type AppHandler struct {
	PathToResources string
	PathToImages    string
	PathToHTML      string

	Logger           Logger
	CharacterService *character.CharacterService
	ItemService      *item.ItemService
}

func (h *AppHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/resource", h.PathToResources)
	router.Static("/image", h.PathToImages)
	router.LoadHTMLGlob(h.PathToHTML)

	router.GET("/", h.HandleIndex)
	router.GET("/eng", h.HandleIndex)
	router.GET("/rus", h.HandleRuIndex)
	router.GET("/por", h.HandlePtIndex)

	router.GET("/items", h.HandleGetItems)
	router.POST("/items", h.HandlePostItems)

	router.GET("/characters", h.HandleGetCharacters)
	router.POST("/characters", h.HandlePostCharacters)

	return router
}
