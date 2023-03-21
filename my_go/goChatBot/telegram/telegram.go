package telegram

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	Token string
	api   *tgbotapi.BotAPI
}

func NewBot(token string) (*TgBot, error) {
	tgBot := &TgBot{
		Token: token,
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 调试模式
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	tgBot.api = bot
	return tgBot, nil
}

func (bot *TgBot) NewChatChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(1)
	u.Timeout = 60

	return bot.api.GetUpdatesChan(u)
}

func (bot *TgBot) SendMsg(chatId int64, msgId int, msgData string) {
	msg := tgbotapi.NewMessage(chatId, msgData)
	msg.ReplyToMessageID = msgId
	// fmt.Printf("wch----- msg: %+v\n", msg)
	bot.api.Send(msg)
}

// 停止服务
func (bot *TgBot) Stop() {
	bot.api.StopReceivingUpdates()
}

// 获取图片数据
func (bot *TgBot) GetImageFile(fileId, filePath string, flage bool) {
	// Download the photo
	file, err := bot.api.GetFile(fileId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// 缓存图片
	if flage {
		filePath := filepath.Join(filePath, fmt.Sprintf("%d.jpg", fileId))
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			if _, err := bot.api.UploadFile(file.FilePath, filePath); err != nil {
				log.Println(err)
				continue
			}
		}
		log.Printf("Downloaded photo to %s", filePath)
	}
}
