package controllers

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// Profile App
type ProfileBot struct {
	bot                  *linebot.Client
	channel_secret       string
	channel_access_token string
}

func Init() (*ProfileBot, error) {
	channel_secret := os.Getenv("CHANNEL_SECRET")
	channel_access_token := os.Getenv("CHANNEL_ACCESS_TOKEN")
	bot, err := linebot.New(channel_secret, channel_access_token)
	if err != nil {
		log.Print("Line bot initiate failed")
		return nil, err
	}

	return &ProfileBot{
		bot:                  bot,
		channel_secret:       channel_secret,
		channel_access_token: channel_access_token,
	}, nil
}
