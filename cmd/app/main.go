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

		ItemService:      item.NewItemService("./item.txt"),
		CharacterService: character.NewCharacterService("./character.txt"),
		Logger:           log.Default(),
	}
	log.Print(handler.CharacterService.GetCharacters())

	handler.InitRoutes().Run()
}
