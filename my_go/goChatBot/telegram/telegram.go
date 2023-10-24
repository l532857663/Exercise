package telegram

import (
	"fmt"
	"goChatBot/utils"
	"io/ioutil"
	"log"
	"net/http"
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
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot.api.GetUpdatesChan(u)
}

func (bot *TgBot) SendMsg(chatId int64, msgId int, msgData string) {
	msg := tgbotapi.NewMessage(chatId, msgData)
	msg.ReplyToMessageID = msgId
	bot.api.Send(msg)
}

func (bot *TgBot) SendPhotoWithFile(chatId int64, msgId int, msgData []byte, caption string) {
	b := tgbotapi.FileBytes{Name: "image.png", Bytes: msgData}
	msg := tgbotapi.NewPhoto(chatId, b)
	msg.Caption = caption
	bot.api.Send(msg)
}

// 停止服务
func (bot *TgBot) Stop() {
	bot.api.StopReceivingUpdates()
}

// 获取图片数据
func (bot *TgBot) GetImageFile(fileId string) (string, error) {
	// Download the photo
	fileURL, err := bot.api.GetFileDirectURL(fileId)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// fmt.Printf("wch---- fileURL: %+v\n", fileURL)
	fBody, err := GetFileBody(fileURL)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// 缓存图片
	filePath, err := DownloadFileToLocal(fBody, fileId)
	if err != nil {
		log.Printf("DownloadFileToLocal error: %+v\n", err)
		return "", err
	}
	log.Printf("Downloaded photo to %s", filePath)
	return filePath, nil
}

// 加载图片内容
func GetFileBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading", url, "-", err)
		return nil, err
	}
	return body, nil
}

// 缓存图片到本地
func DownloadFileToLocal(fBody []byte, fileId string) (string, error) {
	filePath, err := os.Getwd()
	if err != nil {
		log.Println("Getwd error：", err)
		return "", err
	}
	fileName := utils.StringToMD5(fileId)
	filePath = filepath.Join(filePath, fmt.Sprintf("tmpPic/%s.png", fileName))
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer file.Close()
	_, err = file.Write(fBody)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return filePath, nil
}
