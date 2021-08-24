package app

import (
	"github.com/Stezok/game-tools/internal/itemredactor"
	"github.com/gin-gonic/gin"
)

type Logger interface {
	Print(...interface{})
}

type AppHandler struct {
	PathToResources string
	PathToImages    string
	PathToHTML      string

	Logger  Logger
	Service *itemredactor.ItemService
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

	return router
}
