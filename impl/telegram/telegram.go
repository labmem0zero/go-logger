package telegram

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TLogger struct {
	chatId int64
	botAPI *tgbotapi.BotAPI
}

func NewLogger(token string, chatId int64) (l TLogger, err error) {
	l.botAPI, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return
	}
	l.chatId = chatId
	return
}

func (l TLogger) Write(level string, v ...interface{}) {
	v = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.0000"), level}, v...)
	msg := tgbotapi.NewMessage(l.chatId, fmt.Sprintln(v...))
	if _, err := l.botAPI.Send(msg); err != nil {
		fmt.Println(err)
	}
}
