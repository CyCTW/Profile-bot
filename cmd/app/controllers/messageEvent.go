package controllers

import (
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func (app *ProfileBot) MessageEventHandler(event *linebot.Event) {
	// Do something
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		switch message.Text {
		case "你是誰?":
			if err := app.HandleIntroduction(event); err != nil {
				log.Print(err)
			}
		case "你的學歷":
			if err := app.HandleEducation(event); err != nil {
				log.Print(err)
			}
		case "你的經歷":
			if err := app.HandleExperience(event); err != nil {
				log.Print(err)

			}
		case "如何聯絡你?":
			if err := app.HandleContact(event); err != nil {
				log.Print(err)

			}
		case "你的前端專案":
			if err := app.HandleFrontendProject(event); err != nil {
				log.Print(err)

			}
		case "你的後端專案":
			if err := app.HandleBackendProject(event); err != nil {
				log.Print(err)

			}
		case "你的區塊鏈專案":
			if err := app.HandleBlockchainProject(event); err != nil {
				log.Print(err)

			}
		case "你的特質":
			if err := app.HandlePersonal(event); err != nil {
				log.Print(err)

			}
		}
		if _, err := app.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
			log.Print(err)
		}
	default:
		log.Print("Unimplement message event type")
	}
}

func (app *ProfileBot) HandleIntroduction(event *linebot.Event) error {

	flexJsonString := `{
	"type": "bubble",
	"header": {
		"type": "box",
		"layout": "vertical",
		"contents": [
		{
			"type": "text",
			"text": "Software Engineer Intern",
			"weight": "bold"
		}
		]
	},
	"hero": {
		"type": "image",
		"url": "https://lh3.googleusercontent.com/d/1D5O1PUn7Pw4hs7rlHvoHhU14x18Riv23",
		"size": "xl"
	},
	"body": {
		"type": "box",
		"layout": "vertical",
		"contents": [
		{
			"type": "text",
			"text": "Cheng Yuan Chang",
			"size": "xl",
			"weight": "bold",
			"align": "center"
		},
		{
			"type": "text",
			"text": "NTU student",
			"align": "center"
		},
		{
			"type": "separator",
			"margin": "md"
		},
		{
			"type": "box",
			"layout": "vertical",
			"contents": [
			{
				"type": "button",
				"action": {
				"type": "uri",
				"label": "Visit my website👀",
				"uri": "https://cyctw.github.io/"
				},
				"style": "primary"
			},
			{
				"type": "button",
				"action": {
				"type": "message",
				"label": "Contact me✔️",
				"text": "如何聯絡你?"
				},
				"style": "link"
			}
			],
			"paddingTop": "10px"
		}
		]
	},
	"styles": {
		"header": {
		"backgroundColor": "#328fa8"
		}
	}
}`

	contents, err := linebot.UnmarshalFlexMessageJSON([]byte(flexJsonString))
	if err != nil {
		return err
	}
	if _, err := app.bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("My self-introduction", contents)).Do(); err != nil {
		return err
	}
	return nil
}

func (app *ProfileBot) HandleEducation(event *linebot.Event) error {

	flexJsonString := `
	{
		"type": "carousel",
		"contents": [
		  {
			"type": "bubble",
			"header": {
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "text": "National Taiwan University",
				  "align": "center",
				  "weight": "bold"
				}
			  ],
			  "backgroundColor": "#f5a742"
			},
			"hero": {
			  "type": "image",
			  "url": "https://lh3.googleusercontent.com/d/1NgT4DVPmifgPRSGQbaSVBFoPEdoIw-gD",
			  "size": "md",
			  "margin": "none",
			  "offsetTop": "none",
			  "position": "relative",
			  "offsetBottom": "none"
			},
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "text": "Major",
				  "weight": "bold",
				  "size": "xl"
				},
				{
				  "type": "text",
				  "text": "Electrical Engineer (CS)"
				},
				{
				  "type": "separator",
				  "margin": "10px"
				},
				{
				  "type": "text",
				  "text": "Time",
				  "weight": "bold",
				  "size": "xl",
				  "margin": "10px"
				},
				{
				  "type": "text",
				  "text": "2021/09 ~ present"
				},
				{
				  "type": "separator",
				  "margin": "10px"
				},
				{
				  "type": "text",
				  "text": "Research Interest",
				  "margin": "10px",
				  "weight": "bold",
				  "size": "lg"
				},
				{
				  "type": "text",
				  "text": "Blockchain Layer2, Cross chain"
				}
			  ],
			  "offsetTop": "none"
			}
		  },
		  {
			"type": "bubble",
			"header": {
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "text": "National Chiao Tung University",
				  "align": "center",
				  "weight": "bold"
				}
			  ],
			  "backgroundColor": "#3299a8"
			},
			"hero": {
			  "type": "image",
			  "url": "https://lh3.googleusercontent.com/d/1YuQyVXDtsouoj7RvFUCBubApXB2Vdqf8",
			  "size": "md",
			  "offsetTop": "none"
			},
			"body": {
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "text": "Major",
				  "size": "xl",
				  "weight": "bold"
				},
				{
				  "type": "text",
				  "text": "Computer Science"
				},
				{
				  "type": "separator",
				  "margin": "10px"
				},
				{
				  "type": "text",
				  "text": "Time",
				  "size": "xl",
				  "weight": "bold",
				  "margin": "10px"
				},
				{
				  "type": "text",
				  "text": "2017/09 ~ 2021/06"
				}
			  ]
			},
			"styles": {
			  "body": {
				"separator": false
			  }
			}
		  }
		]
	  }`
	contents, err := linebot.UnmarshalFlexMessageJSON([]byte(flexJsonString))
	if err != nil {
		return err
	}
	if _, err := app.bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("My Education", contents)).Do(); err != nil {
		return err
	}
	return nil

}
func (app *ProfileBot) HandleExperience(event *linebot.Event) error {
	return nil

}
func (app *ProfileBot) HandleContact(event *linebot.Event) error {
	messages := linebot.NewTextMessage("請選擇以下一種方式來連繫我!").
		WithQuickReplies(linebot.NewQuickReplyItems(
			linebot.NewQuickReplyButton(
				// app.appBaseURL+"/static/quick/sushi.png",
				"",
				linebot.NewURIAction("Phone", "tel:0987591062")),
			linebot.NewQuickReplyButton(
				// app.appBaseURL+"/static/quick/tempura.png",
				"",
				linebot.NewURIAction("Email", "mailto:0cyctwn@gmail.com")),
			linebot.NewQuickReplyButton(
				"",
				linebot.NewURIAction("Line", "https://line.me/ti/p/WO21tN7ePW")),
			linebot.NewQuickReplyButton(
				"",
				linebot.NewURIAction("Facebook", "https://www.facebook.com/NCTU193/")),
			linebot.NewQuickReplyButton(
				"",
				linebot.NewURIAction("Linkedin", "https://www.linkedin.com/in/cyctw/")),
		))
	_, err := app.bot.ReplyMessage(event.ReplyToken, messages).Do()
	if err != nil {
		return err
	}
	return nil

}
func (app *ProfileBot) HandleFrontendProject(event *linebot.Event) error {
	return nil

}
func (app *ProfileBot) HandleBackendProject(event *linebot.Event) error {
	return nil

}
func (app *ProfileBot) HandleBlockchainProject(event *linebot.Event) error {
	return nil

}
func (app *ProfileBot) HandlePersonal(event *linebot.Event) error {
	return nil

}
