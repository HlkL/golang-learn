package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func helloworld() {
	http.HandleFunc("/", greet)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	post()
}

func get() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		fmt.Printf("r.Method: %v\n", r.Method)
		query := r.URL.Query()
		fmt.Printf("query.Get(\"uid\"): %v\n", query.Get("uid"))
		answer := `{"status": "ok"}`
		w.Write([]byte(answer))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func post()  {
	http.HandleFunc("/post",func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
		// r.ParseForm()
		// fmt.Println(r.PostForm) // 打印form数据
		// fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
		// 2. 请求类型是application/json时从r.Body读取数据
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("read request.Body failed, err:%v\n", err)
			return
		}
		fmt.Println(string(b))
		answer := `{"status": "ok"}`
		w.Write([]byte(answer))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

}