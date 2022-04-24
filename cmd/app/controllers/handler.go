package controllers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type ActivityInput struct {
	Activity string `form:"activity"`
	Date     string `form:"date"`
	Place    string `form:"place"`
}

func getProfile(idToken string) {
	path := "https://api.line.me/oauth2/v2.1/verify"
	postBody := url.Values{}

	postBody.Set("id_token", idToken)
	postBody.Set("client_id", "1657078645")

	client := &http.Client{}
	r, err := http.NewRequest("POST", path, strings.NewReader(postBody.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(r)

	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Print(bodyString)
	log.Println(res)

}

func (app *ProfileBot) ActivityHandler(c *gin.Context) {
	log.Print(c.Request)
	log.Print("Hello!!!!!!!!!!!!!!!!!!!!!")
	// Bind request
	idToken := c.Param("idToken")
	getProfile(idToken)

	var activityInput ActivityInput
	if err := c.ShouldBind(&activityInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	// Create event and store to DB

	c.JSON(200, gin.H{"message": "Hello"})
}

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
