package http

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	httpClient = &http.Client{Timeout: 100 * time.Second}
	httpsInsecureClient = &http.Client{
		Timeout: 100 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
)

func HttpPost(api string, data map[string]interface{}) (code int, body string, err error) {
	req := make(url.Values)
	for key, item := range data {
		req[key] = []string{fmt.Sprintf("%v", item)}
	}

	//把post表单发送给目标服务器
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}
	res, err := client.PostForm(api, req)
	if err != nil {
		return 0, "", err
	}
	defer res.Body.Close()
	bytess, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}
	return res.StatusCode, string(bytess), nil
}

func HttpPostJson(api string, data interface{}) (code int, body string, err error) {
	content, _ := json.Marshal(data)
	return httpPostInner(api, content)
}

func HttpPostStr(api string, data string) (code int, body string, err error) {
	content := []byte(data)
	return httpPostInner(api, content)
}

func httpPostInner(api string, data []byte) (code int, body string, err error) {
	//把post表单发送给目标服务器
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}
	req, err := http.NewRequest("POST", api, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer res.Body.Close()
	bytess, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}
	return res.StatusCode, string(bytess), nil
}

// HttpGet 废弃，没有超时设置
func HttpGet(api string) (code int, body string, err error) {
	res, err := http.Get(api)
	if err != nil {
		return 0, "", err
	}
	defer res.Body.Close()
	bytess, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}
	return res.StatusCode, string(bytess), nil
}

func HttpGetWithHeader(url string, header map[string]string) (code int, body string, err error) {
	// 构造请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, "", errors.New("New req failed!")
	}
	// 添加header
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	// 发送http请求
	response, err := httpClient.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()

	// 读取返回body内容
	bytess, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, "", err
	}
	return response.StatusCode, string(bytess), nil
}

func HttpsGetWithHeaderInsecure(url string, header map[string]string) (code int, body string, err error) {
	// 构造请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, "", errors.New("New req failed!")
	}

	// 添加header
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	// 发送https请求
	response, err := httpsInsecureClient.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()

	// 读取返回body内容
	bytess, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, "", err
	}
	return response.StatusCode, string(bytess), nil
}

func HttpsGetWithHeaderSecure(url string, header map[string]string, rootCA *x509.CertPool) (code int, body string, err error) {
	// 构造请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, "", errors.New("New req failed!")
	}

	// 添加header
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	// 发送https请求
	httpsSecureClient := &http.Client{
		Timeout: 100 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: rootCA},
		},
	}
	response, err := httpsSecureClient.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer response.Body.Close()

	// 读取返回body内容
	bytess, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, "", err
	}
	return response.StatusCode, string(bytess), nil
}

// @Description http的GET请求
// @Author oracle
// @Version 1.0
// @Update 2021/12/06 10:23 init
// @Update 2022/01/04 10:23 取消短连接模式, 抽离chainservice至业务初始和结束来整体控制, 业务流程中复用长链接
func HttpGetWithTimeout(url string, timeout time.Duration) (int, []byte, error) {
	client := &http.Client{
		Timeout:   timeout,
		Transport: &http.Transport{
			// 设置为短连接请求模式
			// DisableKeepAlives: true, // modified @20220104 不再使用短链接的方式
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, body, nil
}

func HttpPostJsonWithHeader(api string, header map[string]string, data interface{}) (code int, body string, err error) {
	content, _ := json.Marshal(data)
	return HttpPostWithHeader(api, header, content)
}

func HttpPostWithHeader(url string, header map[string]string, data []byte) (code int, body string, err error) {
	// 把post表单发送给目标服务器
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	// 添加header
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	res, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer res.Body.Close()
	bytess, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}

	return res.StatusCode, string(bytess), nil
}

func HttpGetWithTimeoutDisableKeepAlives(url string, timeout time.Duration) (int, []byte, error) {
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			// 设置为短连接请求模式
			DisableKeepAlives: true,
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, body, nil
}
