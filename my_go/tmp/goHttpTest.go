package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	HttpAddress   = "http://127.0.0.1:10086/api/bingoo/wallet/yottachain/accounts/v3/create_free"
	HttpAddress1  = "http://127.0.0.1:10086/api/bingoo/wallet/activity/v1/list?entity_id=2354"
	TestNum       = 15
	DefaultClient = &http.Client{
		Timeout: time.Second * 10,
	}
)

type Request struct {
	Headers      map[string]string
	HttpClient   *http.Client
	ResponseTime int64
}

func (r *Request) GetResponseTime(startTime time.Time) {
	start := startTime.UnixNano() / 1e6
	end := time.Now().UnixNano() / 1e6
	r.ResponseTime = end - start
	return
}

func (r *Request) DoHttpReq(method, url string, body io.Reader, index int, wg sync.WaitGroup) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Printf("Platform request error: %+v\n", err)
		return
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	start := time.Now()
	res, err := r.HttpClient.Do(req)
	if err != nil {
		fmt.Printf("http client error: %+v\n", err)
		return
	}
	defer res.Body.Close()
	r.GetResponseTime(start)

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("readAll error: %+v\n", err)
		return
	}

	wg.Done()
	fmt.Printf("wch---------- body %v, time %v ms, index %v\n", string(b), r.ResponseTime, index)
	return
}

func createFreeYottaTest() {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	r := &Request{
		HttpClient: DefaultClient,
		Headers:    headers,
	}
	body := `{
		"language": "CN",
		"entity_id": "2355%d",
		"public_key_owner": "YTA8LJhDgsdncQNzqqSTpypUm698yfcVnVcDpfbAcdTU3TkbGe4FM",
		"public_key_active": "YTA8LJhDgsdncQNzqqSTpypUm698yfcVnVcDpfbAcdTU3TkbGe4FM",
		"description": "ceshishi"
	}`
	var wg sync.WaitGroup

	for i := 0; i < TestNum; i++ {
		wg.Add(1)
		bodyi := fmt.Sprintf(body, i)
		buf := bytes.NewBufferString(bodyi)
		go r.DoHttpReq("POST", HttpAddress, buf, i, wg)
	}
	wg.Wait()
	return
}

func getHttpTest() {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	r := &Request{
		HttpClient: DefaultClient,
		Headers:    headers,
	}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go r.DoHttpReq("GET", HttpAddress1, nil, i, wg)
	}
	wg.Wait()
	return
}

func main() {
	// createFreeYottaTest()
	getHttpTest()
}
