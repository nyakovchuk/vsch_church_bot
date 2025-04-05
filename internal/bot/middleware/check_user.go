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

			fmt.Println("middleware.CheckUser:", user.Username)

			// 1) проверить, что юзер есть в БД
			if notExistsUser(bm, user.ID) {
				tgUserModel := tgUser.ToModel(user)
				bm.Services().User.CreateUser(context.Background(), tgUserModel)
				// 2) если юзера нет в БД, то создать

			}

			return next(c)
		}
	})
}

func notExistsUser(bm BotManager, id int64) bool {
	// возможно стоит возвращать и ошибку,
	// и записать в лог
	return !bm.Services().TgUser.CheckTgId(id)
}
