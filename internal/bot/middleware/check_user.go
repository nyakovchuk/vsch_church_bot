package middleware

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

func CheckUser(bm BotManager) {
	bm.TBot().Use(func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			user := c.Sender()

			fmt.Println("middleware.CheckUser:", user.Username)

			// 1) проверить, что юзер есть в БД

			// 2) если юзера нет в БД, то создать

			return next(c)
		}
	})
}
