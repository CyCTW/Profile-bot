package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {

	r := gin.Default()
	channel_secret := os.Getenv("CHANNEL_SECRET")
	channel_access_token := os.Getenv("CHANNEL_ACCESS_TOKEN")
	bot, err := linebot.New(channel_secret, channel_access_token)
	if err != nil {
		log.Print("Line bot initiate failed")
	}
	r.POST("/callback", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			log.Print("Parse request fail!")
			c.JSON(404, gin.H{"message": "fail"})
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				// Do something
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}

	})
	r.Run()

}
