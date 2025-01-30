package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/tucnak/telebot"
)

const TelegramBotTokenEnv = "TELEGRAM_BOT_TOKEN"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	token := os.Getenv(TelegramBotTokenEnv)
	if token == "" {
		log.Fatal("Токен не найден в переменных окружения")
	}

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
