// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func getRichMenuA() linebot.RichMenu {
	return linebot.RichMenu{
		Size:        linebot.RichMenuSize{Width: 960, Height: 540},
		Selected:    false,
		Name:        "Menu1",
		ChatBarText: "Explore✨",
		Areas: []linebot.AreaDetail{
			{
				Bounds: linebot.RichMenuBounds{X: 27, Y: 101, Width: 150, Height: 150},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "你是誰?",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 303, Y: 271, Width: 150, Height: 150},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "你的學歷",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 631, Y: 198, Width: 150, Height: 150},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "你的經歷",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 820, Y: 23, Width: 120, Height: 77},
				Action: linebot.RichMenuAction{
					Type:            linebot.RichMenuActionTypeRichMenuSwitch,
					RichMenuAliasID: "richmenu-alias-b",
					Data:            "action=richmenu-changed-to-b",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 2, Y: 469, Width: 177, Height: 65},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "如何聯絡你?",
				},
			},
		},
	}
}

func getRichMenuB() linebot.RichMenu {
	return linebot.RichMenu{
		Size:        linebot.RichMenuSize{Width: 960, Height: 540},
		Selected:    false,
		Name:        "Menu2",
		ChatBarText: "Explore✨",
		Areas: []linebot.AreaDetail{
			{
				Bounds: linebot.RichMenuBounds{X: 103, Y: 251, Width: 150, Height: 150},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "你做過什麼專案?",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 382, Y: 3, Width: 150, Height: 150},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "你的前端專案",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 428, Y: 206, Width: 150, Height: 150},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "你的後端專案",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 337, Y: 399, Width: 150, Height: 150},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "你的區塊鏈專案",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 753, Y: 221, Width: 150, Height: 150},
				Action: linebot.RichMenuAction{
					Type: linebot.RichMenuActionTypeMessage,
					Text: "你是個怎樣的人?",
				},
			},
			{
				Bounds: linebot.RichMenuBounds{X: 9, Y: 22, Width: 120, Height: 77},
				Action: linebot.RichMenuAction{
					Type:            linebot.RichMenuActionTypeRichMenuSwitch,
					RichMenuAliasID: "richmenu-alias-a",
					Data:            "action=richmenu-changed-to-a",
				},
			},
		},
	}
}

func main() {
	var (
		mode     = flag.String("mode", "list", "mode of richmenu helper [list|create|link|unlink|bulklink|bulkunlink|get|delete|upload|download|alias_create|alias_get|alias_update|alias_delete|alias_list]")
		aid      = flag.String("aid", "", "alias id")
		uid      = flag.String("uid", "", "user id")
		rid      = flag.String("rid", "", "richmenu id")
		filePath = flag.String("image.path", "", "path to image, used in upload/download mode")
		rType    = flag.String("rType", "", "richmenu type, A or B")
	)
	flag.Parse()
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	switch *mode {
	case "upload":
		if _, err = bot.UploadRichMenuImage(*rid, *filePath).Do(); err != nil {
			log.Fatal(err)
		}
	case "download":
		res, err := bot.DownloadRichMenuImage(*rid).Do()
		if err != nil {
			log.Fatal(err)
		}
		defer res.Content.Close()
		f, err := os.OpenFile(*filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(f, res.Content)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Image is written to %s", *filePath)
	case "alias_create":
		if _, err = bot.CreateRichMenuAlias(*aid, *rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "alias_get":
		if res, err := bot.GetRichMenuAlias(*aid).Do(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%v\n", res)
		}
	case "alias_update":
		if _, err = bot.UpdateRichMenuAlias(*aid, *rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "alias_delete":
		if _, err = bot.DeleteRichMenuAlias(*aid).Do(); err != nil {
			log.Fatal(err)
		}
	case "alias_list":
		res, err := bot.GetRichMenuAliasList().Do()
		if err != nil {
			log.Fatal(err)
		}
		for _, alias := range res {
			log.Printf("%v\n", alias)
		}
	case "link":
		if _, err = bot.LinkUserRichMenu(*uid, *rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "unlink":
		if _, err = bot.UnlinkUserRichMenu(*uid).Do(); err != nil {
			log.Fatal(err)
		}
	case "bulklink":
		if _, err = bot.BulkLinkRichMenu(*rid, *uid).Do(); err != nil {
			log.Fatal(err)
		}
	case "bulkunlink":
		if _, err = bot.BulkUnlinkRichMenu(*uid).Do(); err != nil {
			log.Fatal(err)
		}
	case "list":
		res, err := bot.GetRichMenuList().Do()
		if err != nil {
			log.Fatal(err)
		}
		for _, richmenu := range res {
			log.Printf("%v\n", richmenu)
		}
	case "get_default":
		res, err := bot.GetDefaultRichMenu().Do()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v\n", res)
	case "set_default":
		if _, err = bot.SetDefaultRichMenu(*rid).Do(); err != nil {
			log.Fatal(err)
		}
	case "cancel_default":
		if _, err = bot.CancelDefaultRichMenu().Do(); err != nil {
			log.Fatal(err)
		}
	case "create":
		var richMenu linebot.RichMenu

		switch *rType {
		case "A":
			richMenu = getRichMenuA()
		case "B":
			richMenu = getRichMenuB()
		default:
			log.Fatal("No richmenu type!")
		}

		res, err := bot.CreateRichMenu(richMenu).Do()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.RichMenuID)
	default:
		log.Fatal("implement me")
	}
}

// A: richmenu-5019e12cc320244b5bfb770684499ca5
