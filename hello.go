package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

var helloworld string = initHelloWorld()

func init() {
	println("main package init...")
}

func initHelloWorld() string {
	println("initHelloWorld")
	return "hello,world"
}

// 执行顺序： 全局变量 -> init函数 -> main  优先执行导入包的
// func main() {
// 	println(helloworld)
// 	utils.SayHello()
// }
// 代码执行结果：
// initAge
// initName
// utils package init ...
// initHelloWorld
// main package init...
// hello,world
// hello,go

func main() {
	anonymousFunction()
}

func print() {
	fmt.Println("hello, world")
}

func variable() {
	var number int = 13
	var name string = "hello, go"
	var a, b int = 1, 3
	fmt.Println("number =", number)
	fmt.Println("name =", name)
	fmt.Println(a, b)
	fmt.Println("number byte size: = ", unsafe.Sizeof(number))

	var d int
	var e bool

	fmt.Println(d) // 0
	fmt.Println(e) // false

	lang := "go"
	fmt.Println(lang)

	str := `func Print()  {
    fmt.Println("hello, world")
  	}`

	fmt.Println(str)

	helloworld := "hello" +
		"world"

	fmt.Println(helloworld)
}

func conver() {
	// sprint
	number := 2024
	str := "year:%d"

	rst := fmt.Sprintf(str, number)
	fmt.Print(rst)

	var num int64 = 64
	string := "num to string:%d"
	strnum := strconv.FormatInt(num, 10)
	fmt.Printf("fmt.Sprint(string, strnum): %v\n", fmt.Sprint(string, strnum))

	// float to string
	floatnum := 23.333
	// strconv.FormatFloat(floatnum, 'f', 5, 64): 23.33300
	fmt.Printf("strconv.FormatFloat(floatnum, 'f', 5, 64): %v\n", strconv.FormatFloat(floatnum, 'f', 5, 64))

	// string to basic type
	string2 := "22314"
	// _: ignore result
	intnum, _ := strconv.ParseInt(string2, 0, 64)
	fmt.Println(intnum)

}

func input() {
	var input string
	println("input:")
	fmt.Scanln(&input)
	println("user input :", input)

}

func multipleVarInput() {
	var name string
	var age int
	var pass bool
	println("input:")
	fmt.Scanf("%s %d %t", &name, &age, &pass)
	fmt.Println("name is", name, "age is", age, "pass", pass)
}

func loop() {
	str := "hello, world 你好世界"
	fmt.Println("================")
	// 普通fori循环默认遍历字节，中文会出现乱码
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	println()

	for i := 0; i < 5; i++ {
		fmt.Print(i, "\t")
	}

	fmt.Println("\n================")
	index := 0
	for index < 5 {
		index++
		fmt.Print(index, "\t")
	}

	fmt.Println("\n================")
	count := 5
	for {
		count--

		if count < 1 {
			break
		}

		fmt.Print(count, "\t")
	}

	fmt.Println("\n================")
	// range 按照字符遍历，可以遍历中文，每一个中午占3个字节， 中文也可通过切片遍历
	for index, ch := range str {
		fmt.Printf("index: %d ch: %c\n", index, ch)
	}

	fmt.Println("================")
	// 将字符串转化为切片
	strArr := []rune(str)
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("%c", strArr[i])
	}

	// 函数运行结果：
	// ================
	// hello, world ä½ å¥½ä¸ç
	// 0       1       2       3       4
	// ================
	// 1       2       3       4       5
	// ================
	// 4       3       2       1
	// ================
	// index: 0 ch: h
	// index: 1 ch: e
	// index: 2 ch: l
	// index: 3 ch: l
	// index: 4 ch: o
	// index: 5 ch: ,
	// index: 6 ch:
	// index: 7 ch: w
	// index: 8 ch: o
	// index: 9 ch: r
	// index: 10 ch: l
	// index: 11 ch: d
	// index: 12 ch:
	// index: 13 ch: 你
	// index: 16 ch: 好
	// index: 19 ch: 世
	// index: 22 ch: 界

	// ================
	// hello, world 你好世界%
}

// 匿名函数
func anonymousFunction() {
	a1 := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 12)

	a2 := func(num1 int, num2 int) int {
		return num1 * num2
	}

	println(a1)
	println(a2(2, 4))
}
