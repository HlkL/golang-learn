package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	createLogPrintf()
}

func createLogPrintf() {
	logger := log.New(os.Stdout, "go-log-", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
	// go-log-2024/06/16 15:41:02 main.go:15: 这是自定义的logger记录的日志。
}

func outLog2File()  {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条输出到文件的日志。")
}

func flagLogPrintf()  {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("golang-log")
	log.Println("这是一条很普通的日志。")
	// 2024/06/16 15:35:32.254653 /Users/hougen/code/golang/learn/log/main.go:11: 这是一条很普通的日志。
	//golang-log2024/06/16 15:36:40.563437 /Users/hougen/code/golang/learn/log/main.go:14: 这是一条很普通的日志。
}

func logPrintf() {
	log.Println("这是一条很普通的日志。")
	log.Printf("这是一条%s日志。\n", "普通")
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
	// 2024/06/16 15:24:17 这是一条很普通的日志。
	// 2024/06/16 15:24:17 这是一条普通日志。
	// 2024/06/16 15:24:17 这是一条会触发fatal的日志。
	// exit status 1
}
