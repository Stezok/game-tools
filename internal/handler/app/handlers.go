package app

import (
	"net/http"
	"os"

	"github.com/Stezok/game-tools/internal/handler/app/locale"
	"github.com/Stezok/game-tools/internal/redactor/character"
	"github.com/Stezok/game-tools/internal/redactor/item"
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
	items, err := h.ItemService.GetItems()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		h.Logger.Print(err)
		return
	}
	// h.Logger.Print(items)

	ctx.JSON(http.StatusOK, items)
}

func (h *AppHandler) HandlePostItems(ctx *gin.Context) {
	var items []item.Item
	if err := ctx.BindJSON(&items); err != nil {
		h.Logger.Print(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	file, err := os.Create("ItemInfo.txt")
	if err != nil {
		h.Logger.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = h.ItemService.WriteItems(file, items)
	if err != nil {
		h.Logger.Print(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func (h *AppHandler) HandleGetCharacters(ctx *gin.Context) {
	characters, err := h.CharacterService.GetCharacters()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		h.Logger.Print(err)
		return
	}
	h.Logger.Print(characters[0].Tscsm)

	ctx.JSON(http.StatusOK, characters)
}

func (h *AppHandler) HandlePostCharacters(ctx *gin.Context) {
	var characters []character.Character
	if err := ctx.BindJSON(&characters); err != nil {
		h.Logger.Print(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	file, err := os.Create("CharacterInfo.txt")
	if err != nil {
		h.Logger.Print(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = h.CharacterService.WriteCharacters(file, characters)
	if err != nil {
		h.Logger.Print(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
}
