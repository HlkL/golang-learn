package utils

import "fmt"

// 全局变量，变量或者函数名称首字母大写标识可被其他包所引用，相当于Java中的public
var Age int = initAge()
var Name string = initName()

// init函数在main函数之前执行,且必须没有参数和返回值
func init() {
	println("utils package init ...")
}

func initAge() int {
	println("initAge")
	return 22
}

func initName() string {
	println("initName")
	return "utils"
}

func SayHello() {
	fmt.Println("hello,go")
}
