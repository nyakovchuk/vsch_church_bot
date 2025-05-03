package middleware

import (
	"os"

	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
	"github.com/nyakovchuk/vsch_church_bot/utils"
	"gopkg.in/telebot.v4"
)

func AdminOnly(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if isAdmin(c.Sender().ID) {
			return next(c)
		}

		printer := i18n.Printer(c.Get("lang").(string))

		return c.Send(printer.Sprintf("event.text_location.invalid_coordinates_format"), &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		})
	}
}

func isAdmin(userID int64) bool {
	db_driver := os.Getenv("DB_DRIVER")
	adminID := os.Getenv("TG_ADMIN_ID")
	return db_driver == "sqlite" && utils.Int64ToString(userID) == adminID
}
