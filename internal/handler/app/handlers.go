package app

import (
	"net/http"
	"os"

	"github.com/Stezok/game-tools/internal/handler/app/locale"
	"github.com/Stezok/game-tools/internal/itemredactor"
	"github.com/gin-gonic/gin"
)

func (h *AppHandler) HandleIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", locale.EN_INDEX_LOCALE)
}

func (h *AppHandler) HandleRuIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", locale.RU_INDEX_LOCALE)
}

func (h *AppHandler) HandlePtIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", locale.PT_INDEX_LOCALE)
}

func (h *AppHandler) HandleGetItems(ctx *gin.Context) {
	items, err := h.Service.GetItems()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		h.Logger.Print(err)
		return
	}
	h.Logger.Print(items)

	ctx.JSON(http.StatusOK, items)
}

func (h *AppHandler) HandlePostItems(ctx *gin.Context) {
	var items []itemredactor.Item
	if err := ctx.BindJSON(&items); err != nil {
		h.Logger.Print(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	file, err := os.Create("item.txt")
	if err != nil {
		h.Logger.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = h.Service.WriteItems(file, items)
	if err != nil {
		h.Logger.Print(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
}
