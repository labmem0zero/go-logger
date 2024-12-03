package tlogger

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/labmem0zero/go-logger/formats"
	"github.com/labmem0zero/go-logger/formats/formats_default"
	"github.com/labmem0zero/go-logger/impl"
)

type TLogger struct {
	chatId   int64
	botAPI   *tgbotapi.BotAPI
	format   formats.Format
	settings impl.LoggerSettings
}

func NewLogger(token string, chatId int64, settings impl.LoggerSettings) (l TLogger, err error) {
	l.botAPI, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return
	}
	l.chatId = chatId
	l.format = formats_default.DefaultHTML()
	l.settings = settings
	return
}

func (l TLogger) Write(level string, reqID string, v ...interface{}) {
	if l.settings.Levels != nil {
		if _, ok := l.settings.Levels[level]; !ok {
			return
		}
	}
	v = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.0000"), l.settings.Environment, level, l.settings.AppName, l.settings.AppID, reqID}, v...)
	msg := tgbotapi.NewMessage(l.chatId, l.format.String(v...))
	msg.ParseMode = tgbotapi.ModeHTML
	if _, err := l.botAPI.Send(msg); err != nil {
		fmt.Println(err)
	}
}
