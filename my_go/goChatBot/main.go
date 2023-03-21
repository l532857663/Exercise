package main

import (
	"fmt"
	"goChatBot/initialization"
	"goChatBot/openai"
	"goChatBot/telegram"
	"os"
	"os/signal"
	"syscall"
)

var (
	tgChatId int64 = 5756836802
)

func main() {
	config := initialization.LoadConfig("./config.yaml")
	// 链接TG
	bot, err := telegram.NewBot(config.TokenId)
	if err != nil {
		fmt.Printf("NewBot error: %+v\n", err)
		return
	}

	gpt := openai.NewChatGPT(*config)

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		bot.Stop()
		os.Exit(0)
	}()

	// 创建通道
	updates := bot.NewChatChannel()
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		var (
			user   = update.Message.From.FirstName
			userId = update.Message.From.ID
			chatId = update.Message.MessageID
			// 文本数据
			chatText = update.Message.Text
			// 图片数据
			chatPhoto = update.Message.Photo
			// 文档数据
			chatDoc = update.Message.Document
			// 声音数据
			chatVoice = update.Message.Voice
			// 视频数据
			chatAnima = update.Message.Animation
			chatVideo = update.Message.Video
		)

		fmt.Printf("wch-----[%v] %v\n", userId, user)

		fmt.Printf("wch-----[%v] %v\n", chatId, chatText)
		fmt.Printf("wch-----[%v] %v\n", chatId, chatPhoto)
		fmt.Printf("wch-----[%v] %v\n", chatId, chatAnima)
		fmt.Printf("wch-----[%v] %v\n", chatId, chatVideo)

		fmt.Printf("wch----- data: %+v\n", update.Message.Animation)
		continue

		if chatText != "" {
			resp, err := gpt.SendMsg(chatText)
			if err != nil {
				errStr := fmt.Sprintf("Ask chatGPT failed with error: %v", err)
				bot.SendMsg(update.Message.From.ID, update.Message.MessageID, errStr)
				continue
			}

			bot.SendMsg(update.Message.From.ID, update.Message.MessageID, resp)
		}

		if len(chatPhoto) > 0 {
		}

	}
}
