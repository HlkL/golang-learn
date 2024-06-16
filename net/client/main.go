package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


func main()  {
	sendPostRequest()
}

func sendGetRequest()  {
	resp, err := http.Get("https://www.github.com")
	if err != nil {
		fmt.Println(err)
	}
 	defer resp.Body.Close()
	respBoby := make([]byte, 1024)
	resp.Body.Read(respBoby)
	fmt.Println(string(respBoby))
}

func sendGetRequestWithParam()  {
	param := url.Values{}
	param.Set("uid","13")
	url, err := url.ParseRequestURI("http://127.0.0.1:8080/get")
	if err != nil {
		fmt.Println(err)
	}

	url.RawQuery = param.Encode()
	resp, err := http.Get(url.String())
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(res))
}

func sendPostRequest()  {
	url := "http://127.0.0.1:8080/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=张三&age=18"
	contentType := "application/json"
	data := `{"name":"张三","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}