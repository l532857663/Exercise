package openai

import (
	"goChatBot/utils"
	"log"
)

func (gpt ChatGPT) SendMsg(req string) (string, error) {
	msgs := []Messages{
		{Role: "user", Content: req},
	}
	resp, err := gpt.Completions(msgs)
	if err != nil {
		return "", err
	}
	return resp.Content, nil
}

func (gpt ChatGPT) ImageGenerate(prompt, size string) ([]byte, error) {
	imageBs64, err := gpt.GenerateOneImage(prompt, size)
	if err != nil {
		log.Printf("GenerateOneImage failed with error: %v", err)
		return nil, err
	}
	return utils.Base64ToByte(imageBs64)
}

func (gpt ChatGPT) ImageVariantion(image, size string) ([]byte, error) {
	ConvertToRGBA(image, image)
	err := VerifyPngs([]string{image})
	if err != nil {
		log.Printf("VerifyPngs failed with error: %v", err)
		return nil, err
	}
	imageBs64, err := gpt.GenerateOneImageVariation(image, size)
	if err != nil {
		log.Printf("GenerateOneImageVariation failed with error: %v", err)
		return nil, err
	}
	return utils.Base64ToByte(imageBs64)
}
