package controllers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func verifyLineSignature(req *http.Request, channel_secret string) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print("Read error")
	}

	decoded, err := base64.StdEncoding.DecodeString(req.Header.Get("x-line-signature"))
	if err != nil {
		log.Print("decoded error")
	}
	hash := hmac.New(sha256.New, []byte(channel_secret))
	hash.Write(body)

	res := hmac.Equal(hash.Sum(nil), decoded)
	if res != true {
		return errors.New("Verification fail")
	} else {
		log.Print("Verify success!")
	}

	// Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return nil
}

func (app *ProfileBot) CallbackHandler(c *gin.Context) {

	// Verification header
	err := verifyLineSignature(c.Request, app.channel_secret)
	if err != nil {
		log.Print("Verify Failed")
		c.JSON(404, gin.H{"message": "Header verification fail"})
		return
	}

	events, err := app.bot.ParseRequest(c.Request)
	if err != nil {
		log.Print("Parse request fail!")
		c.JSON(404, gin.H{"message": "Parse request fail"})
		return
	}

	for _, event := range events {
		log.Printf("Got event %v", event)
		switch event.Type {
		case linebot.EventTypeMessage:
			app.MessageEventHandler(event)
		case linebot.EventTypeFollow:
			log.Print("Follow")
		case linebot.EventTypeUnfollow:
			log.Print("Unfollow")
		case linebot.EventTypePostback:
			log.Print("Postback event")
		default:
			log.Printf("Unknown event %v", event)
		}

	}

}
