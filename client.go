package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	REQ_URL = "http://httpbin.org"
)

func main() {
	values := url.Values{}
	values.Add("param1", "値")
	fmt.Println(values.Encode())
	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	get(client, values)
}

func get(client *http.Client, values url.Values) {
	req, err := http.NewRequest("GET", REQ_URL+"/get", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.URL.RawQuery = values.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	execute(resp)
}

func execute(resp *http.Response) {
	// response bodyを文字列で取得するサンプル
	// ioutil.ReadAllを使う
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(b))
	}
}
