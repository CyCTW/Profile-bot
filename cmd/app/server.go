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
	r.Static("assets", "./static/line-liff-v2-starter/src/nextjs/out")
	r.POST("/callback", app.CallbackHandler)
	r.POST("/activity/:idToken", app.ActivityHandler)
	r.Run()

}
