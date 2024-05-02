package main

import (
	"fmt"
	"strings"
)

func main() {
	// 闭包产生内存逃逸
	fun := addUpper()
	fmt.Printf("addUpper(): %v\n", fun())
	fmt.Printf("addUpper(): %v\n", fun())
	fmt.Printf("addUpper(): %v\n", fun())
	// 运行结果
	// addUpper(): 1
	// addUpper(): 2
	// addUpper(): 3

	fmt.Println("================")
	fun2 := makeSuffix(".go")
	fmt.Printf("fun2(\"main.go\"): %v\n", fun2("main.go"))
	fmt.Printf("fun2(\"main.go\"): %v\n", fun2("utils"))
	fmt.Printf("fun2(\"main.go\"): %v\n", fun2("hello"))
	// fun2("main.go"): main.go
	// fun2("main.go"): utils.go
	// fun2("main.go"): hello.go

}

func addUpper() func() int {
	var num int

	// 匿名函数引用外部变量
	return func() int {
		num++
		return num
	}

}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
}
