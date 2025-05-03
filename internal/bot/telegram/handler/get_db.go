package handler

import (
	"os"
	"time"

	"gopkg.in/telebot.v4"
)

const messageFileNotFound = "Файл db.sqlite не найден."

var timeFormat = "2006-01-02_15-04"

func HandleGetDB(bm BotManager) func(telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		filePath := os.Getenv("DSN")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return c.Send(messageFileNotFound)
		}

		return c.Send(&telebot.Document{
			File:     telebot.FromDisk(filePath),
			FileName: getFileName(),
		})
	}
}

func getFileName() string {
	currentDate := time.Now().Format(timeFormat)

	return "db-" + currentDate + ".sqlite"
}
