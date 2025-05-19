package middleware

import (
	"context"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/user"
	"github.com/nyakovchuk/vsch_church_bot/utils"
	"gopkg.in/telebot.v4"
)

func CheckUser(bm BotManager) {
	bm.TBot().Use(func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			tgUser := c.Sender()

			if unregisteredUser(bm, tgUser.ID) {
				userModel := user.FromTelebotUser(tgUser)
				platformId := bm.SharedData().Platform.ID

				err := bm.Services().User.Register(context.Background(), platformId, userModel)
				if err != nil {
					bm.LoggerError(c, err)
				}
				bm.LoggerMessage(c, "User registered")
			}

			return next(c)
		}
	})
}

func unregisteredUser(bm BotManager, id int64) bool {
	platformId := bm.SharedData().Platform.ID
	externalId := utils.Int64ToString(id)

	exists, err := bm.Services().User.IsRegistered(platformId, externalId)
	if err != nil {
		// bm.LoggerInfo()
		return false
	}
	return !exists
}
