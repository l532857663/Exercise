package global

import (
	"fmt"
	"goChatBot/initialization"
	"goChatBot/openai"
	"goChatBot/telegram"
)

var (
	gpt *openai.ChatGPT
	bot *telegram.TgBot
)

func Init() error {
	config := initialization.LoadConfig("./config.yaml")
	// 链接TG
	var err error
	bot, err = telegram.NewBot(config.TokenId)
	if err != nil {
		fmt.Printf("NewBot error: %+v\n", err)
		return err
	}
	gpt = openai.NewChatGPT(*config)
	return nil
}

func Stop() {
	bot.Stop()
}

func CreateBotReceive() {
	// 创建通道
	updates := bot.NewChatChannel()
	var (
		userBotType map[int64]string
	)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		var (
			msgInfo = update.Message
			userId  = update.Message.From.ID
			chatId  = update.Message.Chat.ID
			msgId   = update.Message.MessageID

			flage string
			text  string
		)
		switch update.Message.Command() {
		case "help":
			text = "Send a message to start talking with ChatGPT. You can use /reload at any point to clear the conversation history and start from scratch (don't worry, it won't delete the Telegram messages)."
		case "start":
			text = "Send a message to start talking with ChatGPT. You can use /reload at any point to clear the conversation history and start from scratch (don't worry, it won't delete the Telegram messages)."
		case "reload":
			text = "Started a new conversation. Enjoy!"
			flage = ""
		case AskTypeChat:
			text = "Send a message to talking with ChatGPT."
			userBotType[userId] = AskTypeChat
		case AskTypeImage:
			text = "The next step is to create pictures.Please send description information"
			userBotType[userId] = AskTypeImage
		}

		userFlage, ok := userBotType[userId]
		if !ok {
			userFlage = flage
		}

		if text != "" {
			bot.SendMsg(chatId, msgId, text)
			text = ""
			continue
		}

		go DealWithMsg(msgInfo, userFlage)
	}
}
