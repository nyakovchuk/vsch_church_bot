package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tucnak/telebot"
)

func main() {
	// Вставьте ваш токен, полученный от @BotFather
	token := "7895240847:AAH9y2APPigER0JaHNlyz32V-uVw4SvwHuE"

	// Создаем нового бота
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	// Обработка команды /start
	bot.Handle("/start", func(m *telebot.Message) {
		bot.Send(m.Chat, "Привет! Я ваш Telegram-бот.")
	})

	// Обработка команды /help
	bot.Handle("/help", func(m *telebot.Message) {
		bot.Send(m.Chat, "Я могу помочь вам с различными задачами. Введите команду, чтобы узнать, что я могу.")
	})

	// Запуск бота
	fmt.Println("Бот запущен...")
	bot.Start()
}
