package main

import "github.com/Stezok/game-tools/internal/handler/app"

func main() {

	handler := &app.AppHandler{
		PathToResources: "../../web",
		PathToImages:    "../../assets",
		PathToHTML:      "../../web/html/*.html",
	}

	handler.InitRoutes().Run()
}
