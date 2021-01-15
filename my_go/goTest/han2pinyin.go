package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	/*
		str, err := pinyin.New("hi,我是中国人").Split("").Mode(pinyin.InitialsInCapitals).Convert()
		if err != nil {
			fmt.Println("error:", err)
		}
	*/
	phonecode, smsTemplateID, smsSign, randnum, smsSdkAppid := "1", "2", "3", "4", "5"
	params := "{\"PhoneNumberSet\":[\"" + phonecode + "\"],\"TemplateID\":" +
		"\"" + smsTemplateID + "\",\"Sign\":\"" + smsSign + "\",\"TemplateParamSet\":[\"" +
		randnum + "\"],\"SmsSdkAppid\":\"" + smsSdkAppid + "\"}"
	params1 := fmt.Sprintf(`{"PhoneNumberSet":["%s"],"TemplateID":"%s","Sign":"%s","TemplateParamSes":["%s"],"SmsSdkAppid":"%s"}`, phonecode, smsTemplateID, smsSign, randnum, smsSdkAppid)
	fmt.Println("str: ", params)
	fmt.Println("str1:", params1)
}
