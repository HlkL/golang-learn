package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	multipleVarInput()
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
