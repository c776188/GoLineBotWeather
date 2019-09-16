package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/line/line-bot-sdk-go/linebot"
)

type LineCallback struct {
	beego.Controller
}

func (c *LineCallback) Post() {
	bot, err := linebot.New(
		beego.AppConfig.String("LINE_CHANNEL_SECURITY"),
		beego.AppConfig.String("LINE_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(c.Ctx.Request)
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			// userID := event.Source.UserID
			// groupID := event.Source.GroupID
			// RoomID := event.Source.RoomID
			// replyToken := event.ReplyToken

			leftBtn := linebot.NewMessageAction("left", "left clicked")
			rightBtn := linebot.NewMessageAction("right", "right clicked")

			template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)

			message1 := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)

			// _, err := bot.PushMessage(userID, message...).Do()
			// if err != nil {
			// 	// Do something when some bad happened
			// }

			// linebot.NewTextMessage(message.Text)
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, message1).Do(); err != nil {
					log.Print(message)
					log.Print(err)
				}
			}
		}
	}
}
