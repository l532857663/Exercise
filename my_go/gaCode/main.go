package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/pquerna/otp/totp"
)

const (
	recaptchaURL    = "https://www.google.com/recaptcha/api/siteverify"
	recaptchaSecret = "E7VYW4NNXM3Y33VTUCA4JODOO2O3NELO"
	account         = "w@163.com"
	imgSrc          = "https://t7.baidu.com/it/u=1595072465,3644073269&fm=193&f=GIF"
	result          = `
<div>key:<br/>%s</div><div>QrCode:<br/><div id="qrcode"></div></div>
<script src="https://cdn.jsdelivr.net/npm/jquery/dist/jquery.min.js"></script>
<script src="https://cdn.rawgit.com/davidshimjs/qrcodejs/gh-pages/qrcode.min.js"></script>
<script>
    // 使用 JavaScript 生成二维码
    var qrcode = new QRCode(document.getElementById("qrcode"), {
        text: "%s",  // 替换为你的二维码内容
        width: 128,
        height: 128
    });
</script>
`
)

func main() {
	http.HandleFunc("/generate", generateHandler)
	http.HandleFunc("/verify", verifyHandler)
	log.Fatal(http.ListenAndServe(":4396", nil))
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key, qrCode, err := generateRecaptcha(account)
	if err != nil {
		http.Error(w, "generateRecaptcha", http.StatusInternalServerError)
		return
	}

	body := fmt.Sprintf(result, key, qrCode)
	fmt.Fprintln(w, body)
}

func generateRecaptcha(account string) (string, string, error) {
	// 生成密钥
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: account,
	})
	if err != nil {
		return "", "", err
	}
	fmt.Printf("generateRecaptcha: key[%s], qrCode[%s]\n", key.Secret(), key.URL())
	return key.Secret(), key.URL(), nil
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取 POST 请求中的验证码响应和用户 IP 地址
	response := r.FormValue("g-recaptcha-response")
	remoteIP := r.RemoteAddr
	fmt.Printf("wch---- response: %+v\n", response)
	fmt.Printf("wch---- remoteIP: %+v\n", remoteIP)

	if totp.Validate(response, recaptchaSecret) {
		fmt.Printf("wch------ dddddd\n")
	}
	// 验证验证码
	if verifyRecaptcha(response, remoteIP) {
		fmt.Fprintln(w, "Verification successful")
	} else {
		fmt.Fprintln(w, "Verification failed")
	}
}

func verifyRecaptcha(response, remoteIP string) bool {
	// 构建 POST 请求参数
	data := url.Values{}
	data.Set("secret", recaptchaSecret)
	data.Set("response", response)
	data.Set("remoteip", remoteIP)

	// 发送 POST 请求到 Google reCAPTCHA 验证 API
	resp, err := http.PostForm(recaptchaURL, data)
	if err != nil {
		log.Println("Failed to verify reCAPTCHA:", err)
		return false
	}
	defer resp.Body.Close()

	// 读取并解析响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response:", err)
		return false
	}
	fmt.Printf("wch------ body: %s", string(body))

	// 检查响应是否包含 "success": true
	if strings.Contains(string(body), "\"success\": true") {
		return true
	}

	return false
}
