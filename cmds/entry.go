package cmds

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/sirupsen/logrus"
)

var (
	reBing    *regexp.Regexp = regexp.MustCompile(`(?i)^\/bing-[\w/:.\p{Han}]{1,20}$`)
	reWnacg   *regexp.Regexp = regexp.MustCompile(`(?i)^w\d{1,5}$`)
	reNhentai *regexp.Regexp = regexp.MustCompile(`(?i)^n\d{1,6}$`)
)

func TextMessageEntryPoint(bot *messaging_api.MessagingApiAPI, e webhook.MessageEvent, message webhook.TextMessageContent) {
	var messages []messaging_api.MessageInterface

	message.Text = strings.TrimSpace(message.Text)

	if message.Text == "/test" {
		messages = append(messages, messaging_api.TextMessage{
			Text: "✅messaging_api 測試成功",
		})
	}

	if message.Text == "/抽" || message.Text == "/抽籤" {
		messages = append(messages, messaging_api.TextMessage{
			Text: fmt.Sprintf("您抽到的數字為: %s", draw()),
		})
	}

	if reBing.MatchString(message.Text) {
		retrunString, err := bingSearch(message.Text[8:])
		if err != nil {
			logrus.Errorf("bing搜尋功能發生錯誤: %s", err)
			logrus.Error(fmt.Sprintf("bing搜尋功能發生錯誤: %s", err))
		} else {
			messages = append(messages, messaging_api.TextMessage{
				Text: retrunString,
			})
		}
	}

	if strings.Contains(message.Text, "查") {
		messages = append(messages, messaging_api.TextMessage{
			Text: "以下是查分器連結 請妥善使用～\nhttps://redive.estertion.win/arcaea/probe/",
		})
	}

	if strings.Contains(strings.ToLower(message.Text), "vc") {
		messages = append(messages, messaging_api.TextMessage{
			Text: "仙草快跟他結婚#",
		})
	}

	if strings.Contains(message.Text, "天堂門") {
		messages = append(messages, messaging_api.TextMessage{
			Text: "Snowth快去P!!!",
		})
	}

	if strings.Contains(strings.ToLower(message.Text), "運勢") || strings.ContainsAny(strings.ToLower(message.Text), "運勢") {
		messages = append(messages, messaging_api.TextMessage{
			Text: fmt.Sprintf("💫您今天的運勢: %s💫", fortunate()),
		})
	}

	if reNhentai.MatchString(message.Text) {
		messages = append(messages, messaging_api.TextMessage{
			Text: fmt.Sprint("https://nhentai.net/g/" + message.Text[1:]),
		})
	}

	if reWnacg.MatchString(message.Text) {
		returnString, err := wnacgCheck(message.Text[1:])
		if err != nil {
			logrus.Errorf("wnacg功能發生錯誤: %s", err)
			logrus.Error(fmt.Sprintf("wnacg功能發生錯誤: %s", err))
		}
		messages = append(messages, messaging_api.TextMessage{
			Text: returnString,
		})
	}

	if strings.Contains(strings.ToLower(message.Text), "ciallo") ||
		strings.Contains(strings.ToLower(message.Text), "(∠・ω< )") ||
		strings.Contains(strings.ToLower(message.Text), "恰囉") ||
		strings.Contains(strings.ToLower(message.Text), "洽囉") {
		messages = append(messages, messaging_api.TextMessage{
			Text: "Ciallo～(∠・ω< )⌒★",
		})
	}

	if len(messages) != 0 {
		_, err := bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages:   messages,
			},
		)
		if err != nil {
			logrus.Error(err)
		} else {
			logrus.Info(fmt.Sprintf("使用者說: %s", message.Text))
		}
	} else {
		logrus.Info(fmt.Sprintf("使用者說: %s", message.Text))
	}

}
