package main

import (
	"log"

	"github.com/cyctw/line-profile-bot/cmd/app/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	app, err := controllers.Init()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.POST("/callback", app.CallbackHandler)
	r.Run()

}
