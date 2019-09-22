package main

import (
	"log"
	"net/http"
	
	"github.com/line/line-bot-sdk-go/linebot"
)
func main() {
	bot, err := linebot.New(
		"584c8dc1777c8a85762c2cfc429fff01",
		"CZObiwkh2GJNt2zq/dlIM0oqaY45BfIWbe+XFS5KGQy6gNrZ1qti2IGUKj7kDa5Rubjqt2fZ9d9U25a/WLCySfvK52bOOnKKkEGHk6qncFeoQulYcAjb+fqqdt1L/Ba8hSTjCvTH3K+tcvG9TiW9zAdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}