package main

import (
	"help/internal/api"
	"help/internal/domain/AnekdotProviders"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	go e.Start(":8080")

	anekdotProvider := AnekdotProviders.New(
		&http.Client{},
	)
	apiDomain := api.New(anekdotProvider)
	e.GET(api.GetAnekdotHandlerName, apiDomain.GetAnekdotHandler)
	apiDomain.StartBotHandler(anekdotProvider)
}
