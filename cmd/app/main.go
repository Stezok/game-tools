package main

import (
	"log"

	"github.com/Stezok/game-tools/internal/handler/app"
	"github.com/Stezok/game-tools/internal/itemredactor"
)

func main() {

	handler := &app.AppHandler{
		PathToResources: "../../web",
		PathToImages:    "../../assets",
		PathToHTML:      "../../web/html/*.html",

		Service: itemredactor.NewItemService("./item.txt"),
		Logger:  log.Default(),
	}

	handler.InitRoutes().Run()
}
