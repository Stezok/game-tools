package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AppHandler) HandleIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
