package handler

import (
	"os"
	"time"

	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
	"github.com/nyakovchuk/vsch_church_bot/utils"
	"gopkg.in/telebot.v4"
)

func HandleGetDB(bm BotManager) func(telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		db_driver := os.Getenv("DB_DRIVER")
		adminID := os.Getenv("TG_ADMIN_ID")
		if db_driver == "sqlite" && utils.Int64ToString(c.Sender().ID) == adminID {
			filePath := os.Getenv("DSN")
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				return c.Send("Файл db.sqlite не найден.")
			}

			currentDate := time.Now().Format("2006-01-02_15-04")
			fileName := "db-" + currentDate + ".sqlite"

			return c.Send(&telebot.Document{
				File:     telebot.FromDisk(filePath),
				FileName: fileName,
			})
		}

		printer := i18n.Printer(c.Get("lang").(string))

		return c.Send(printer.Sprintf("event.text_location.invalid_coordinates_format"), &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		})
	}
}
