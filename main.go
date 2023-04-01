package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kappaprideonly/ege_bot_2.0/database"
)

func main() {
	key, exist := os.LookupEnv("KEY_BOT")
	log.Printf("%s\n", key)
	if exist == false {
		log.Panic("Key doesn't exist")
	}
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}
	database.Init()
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			log.Printf("%d", update.Message.From.ID)
			id := update.Message.From.ID
			name := update.Message.From.FirstName
			result, _ := database.FindUser(uint(id))
			if result.Error != nil {
				log.Printf("Can't find user with id=%d", id)
				database.CreateUser(uint(id), name, 0)
			} else {
				log.Printf("User find!")
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Вы авторизованы!"))
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
