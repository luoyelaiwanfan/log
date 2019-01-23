package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Test struct {
	A int
	B int
}


func get() {
	//client := &http.Client{Timeout: time.Duration(time.Duration(1) * time.Second)}
	client := &http.Client{}
	url := "http://127.0.0.1:12345/heartbeat"

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("err:", err)
	}

	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer response.Body.Close()
	if response == nil {
		fmt.Println("fail")
	}

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)

	////返回的状态码
	//status := response.StatusCode
	//
	//fmt.Println(status)
}

func post() {
	test := Test{1, 2}

	testBy, _ := json.Marshal(test)

	bytes.NewReader(testBy)

	req, err := http.NewRequest("POST","http://127.0.0.1:12345/reload", bytes.NewReader(testBy))
	if err != nil {
		// handle error
	}

	// 表单方式(必须)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//AJAX 方式请求
	//req.Header.Add("x-requested-with", "XMLHttpRequest")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println("body:", string(body))
	//返回的状态码
	status := resp.StatusCode

	fmt.Println(status)
}
func main() {
	post()
	t := time.NewTicker(time.Second*2)

	for {
		select {
		case <-t.C:
			go get()

			//case <-ticker.C:
			//	fmt.Println("break heart")
		}
	}

	//get()


	var sig = make(chan os.Signal, 1)

	for {
		select {
		case <-sig:
			fmt.Println("Signal received.  Shutting down...\n")
			return
			//case <-ticker.C:
			//	fmt.Println("break heart")
		}
	}
}