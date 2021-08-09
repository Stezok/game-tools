package app

import "github.com/gin-gonic/gin"

type AppHandler struct {
	PathToResources string
	PathToImages    string
	PathToHTML      string
}

func (h *AppHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/resource", h.PathToResources)
	router.Static("/image", h.PathToImages)
	router.LoadHTMLGlob(h.PathToHTML)

	router.GET("/", h.HandleIndex)

	return router
}
