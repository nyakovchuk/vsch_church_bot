package handler

import (
	"os"
	"time"

	"gopkg.in/telebot.v4"
)

const (
	timeFormat                   = "2006-01-02_15-04"
	messageFileNotFound          = "Файл базы данных не найден."
	messageFileNotFoundForSqlite = "База данных не sqlite. Файл базы данных не может быть переслан."
)

func HandleGetDB(bm BotManager) func(telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		db_driver := os.Getenv("DB_DRIVER")
		if db_driver == "sqlite" {

			filePath := os.Getenv("DSN")
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				return c.Send(messageFileNotFound)
			}

			return c.Send(&telebot.Document{
				File:     telebot.FromDisk(filePath),
				FileName: getFileName(),
			})
		}

		return c.Send(messageFileNotFoundForSqlite)
	}
}

func getFileName() string {
	currentDate := time.Now().Format(timeFormat)

	return "db-" + currentDate + ".sqlite"
}
