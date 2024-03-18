package TelegramBot

// type GetTelegramUser interface {
// 	GetTelegramUser(userID int) error
// }
type SendJoke interface {
	sendJoke(userId int, joke string) error
}
