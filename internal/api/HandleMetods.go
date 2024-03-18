package api

import (
	"context"
	"help/internal/domain/AnekdotProviders"
	"help/internal/domain/TelegramBot"
	"help/internal/dto"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const GetAnekdotHandlerName = "/anekdot"

func (d *Domain) GetAnekdotHandler(c echo.Context) error {
	AnekdotProviders := &AnekdotProviders.Domain{}
	anekdot, err := d.anekdotProvider.GetAnekdot(c.Request().Context(), AnekdotProviders.GenerateURL())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			ErrorMessage: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.Anekdot{
		Text: anekdot,
	})
}
func (d *Domain) StartBotHandler(AnekdotProviders *AnekdotProviders.Domain) {

	TgBotDomain := TelegramBot.New()

	anekdot, err := d.anekdotProvider.GetAnekdot(context.Background(), AnekdotProviders.GenerateURL())
	if err != nil {
		log.Println("Ошибка при получении анекдота:", err)
		return
	}

	TgBotDomain.GetTelegramUser(anekdot)
}
