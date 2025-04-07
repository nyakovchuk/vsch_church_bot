package middleware

import (
	"context"
	"fmt"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/tgUser"
	"gopkg.in/telebot.v4"
)

func CheckUser(bm BotManager) {
	bm.TBot().Use(func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			user := c.Sender()

			if unregisteredUser(bm, user.ID) {
				tgUserModel := tgUser.ToModel(user)
				err := bm.Services().User.Register(context.Background(), tgUserModel)
				if err != nil {
					fmt.Println("user check:", err)
					// залогировать
				}
			}

			return next(c)
		}
	})
}

func unregisteredUser(bm BotManager, id int64) bool {
	exists, err := bm.Services().TgUser.CheckTgId(id)
	if err != nil {
		// bm.LoggerInfo()
		return false
	}
	return !exists
}
