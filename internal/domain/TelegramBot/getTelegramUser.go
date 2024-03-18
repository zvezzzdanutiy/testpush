package TelegramBot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (d *Domain) GetTelegramUser(anekdot string) {
	var userID int
	var token string = "7122476551:AAGRldhloWEs-_jWsEkOTMZEsXhGE0dbXWQ"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать! Чтобы получить анекдоты, введите /generate")
			userID = update.Message.From.ID
			d.ReturnID(userID)
			bot.Send(msg)
		} else if update.Message.Text == "/generate" {
			go func(update tgbotapi.Update) {
				msg := tgbotapi.NewMessage(int64(userID), anekdot)
				bot.Send(msg)
			}(update)
		}
	}

}
func (d *Domain) ReturnID(id int) int {
	return id
}
