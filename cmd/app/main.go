package main

import (
	"log"

	"github.com/Stezok/game-tools/internal/handler/app"
	"github.com/Stezok/game-tools/internal/redactor/character"
	"github.com/Stezok/game-tools/internal/redactor/item"
)

func main() {

	handler := &app.AppHandler{
		PathToResources: "../../web",
		PathToImages:    "../../assets",
		PathToHTML:      "../../web/html/*.html",

		ItemService:      item.NewItemService("./ItemInfo.txt"),
		CharacterService: character.NewCharacterService("./CharacterInfo.txt"),
		Logger:           log.Default(),
	}

	handler.InitRoutes().Run()
}
