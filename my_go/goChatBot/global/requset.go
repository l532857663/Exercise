package global

import (
	"fmt"
	"goChatBot/utils"
	"log"
	"runtime/debug"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	AskTypeChat  = "chat"
	AskTypeImage = "image"
	BotName      = "askBJJ_bot"
)

func DealWithMsg(msg *tgbotapi.Message, flage string) {
	var (
		user   = msg.From.UserName
		chatId = msg.Chat.ID
		msgId  = msg.MessageID
		// 文本数据
		chatText = utils.DeleteStringBot(msg.Text, "@"+BotName)
		// 图片数据
		chatPhoto = msg.Photo
		// // 文档数据
		// chatDoc = update.Message.Document
		// // 声音数据
		// chatVoice = update.Message.Voice
		// // 视频数据
		// chatAnima = update.Message.Animation
		// chatVideo = update.Message.Video
		errMsg string
	)
	defer func() {
		if err := recover(); err != nil {
			// 打印崩溃信息
			fmt.Println("程序崩溃：", err)
			// 打印堆栈跟踪信息
			fmt.Println(string(debug.Stack()))

			errMsg = fmt.Sprintf("Ask chatGPT failed with error: %v", err)
		}
		if errMsg != "" {
			bot.SendMsg(chatId, msgId, errMsg)
		}
	}()

	fmt.Printf("wch-----[%v] %v, msgId:%+v\n", chatId, user, msgId)
	switch flage {
	case AskTypeImage:
		size := "512x512"
		if chatText != "" {
			msg, err := gpt.ImageGenerate(chatText, size)
			if err != nil {
				errMsg = fmt.Sprintf("Ask chatGPT image variantion error: %v\n", err)
				log.Println(errMsg)
				return
			}
			bot.SendPhotoWithFile(chatId, msgId, msg, chatText)
			return
		}
		if len(chatPhoto) > 0 {
			image, err := bot.GetImageFile(chatPhoto[len(chatPhoto)-1].FileID)
			if err != nil {
				errMsg = fmt.Sprintf("Get chat photo error: %v\n", err)
				log.Println(errMsg)
				return
			}
			msg, err := gpt.ImageVariantion(image, size)
			if err != nil {
				errMsg = fmt.Sprintf("Ask chatGPT image variantion error: %v\n", err)
				log.Println(errMsg)
				return
			}
			bot.SendPhotoWithFile(chatId, msgId, msg, "")
			return
		}
	case AskTypeChat, "":
		if chatText == "" {
			errMsg = fmt.Sprintf("Not get chat infomation!")
			log.Println(errMsg)
			return
		}
		resp, err := gpt.SendMsg(chatText)
		if err != nil {
			errMsg = fmt.Sprintf("Send msg error: %+v\n", err)
			log.Println(errMsg)
			return
		}
		bot.SendMsg(chatId, msgId, resp)
		return
	}
	errMsg = fmt.Sprintf("Not get chat infomation!")
}
