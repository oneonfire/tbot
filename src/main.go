package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
)

const (
	Token  = "1195820277:AAEc4fmOA3ARtAP0MFn-ucKPCpH3lVSdyMI"
	ChatID = 253517276

//       Version = 0.2
)

func main() {
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start":
				msg.Text = "Цей бот підтримує команди: /help, /version, /sayhi, /status, /withArgument, /html"
			case "help":
				msg.Text = "введи /sayhi чи /status."
			case "version":
				msg.Text = "tbot 0.2 created with go"
			case "sayhi":
				msg.Text = "Привіт :)"
			case "status":
				msg.Text = "В мене все добре."
			case "withArgument":
				msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
			case "html":
				msg.ParseMode = "html"
				msg.Text = "This will be interpreted as HTML, click <a href=\"https://www.example.com\">here</a>"
			case "image":
				data, _ := ioutil.ReadFile("image.png")
				b := tgbotapi.FileBytes{Name: "image.jpg", Bytes: data}
				msg := tgbotapi.NewPhotoUpload(ChatID, b)
				msg.Caption = "Test"
				bot.Send(msg)
			case "file":
				data, _ := ioutil.ReadFile("Luchkov.mp4")
				b := tgbotapi.FileBytes{Name: "Luchkov.mp4", Bytes: data}
				msg := tgbotapi.NewDocumentUpload(ChatID, b)
				bot.Send(msg)
			default:
				msg.Text = "Я не знаю цієї команди"
			}
			bot.Send(msg)
		}
	}
}
