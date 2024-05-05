package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var name string
	var pwd string
	var port int
	flag.StringVar(&name, "u", "", "用户名为空")
	flag.StringVar(&pwd, "pwd", "", "密码为空")
	flag.IntVar(&port, "p", 0, "端口")

	flag.Parse()

	fmt.Printf("用户名: %s,用户密码: %s,端口号: %d\n", name, pwd, port)
}

func args() {
	args := os.Args
	for _, str := range args {
		fmt.Println(str)
	}
}
