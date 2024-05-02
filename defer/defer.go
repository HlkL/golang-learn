package main

import "fmt"

func main() {
	sum := sum(10, 12)
	fmt.Printf("sum(10, 12): %v\n", sum)
	// 运行结果
	// sum: 22,num1: 11,num2: 13
	// num2:  12
	// num1:  10
	// sum(10, 12): 22
}

func sum(num1 int, num2 int) int {
	// defer 会将标记的语句，压入一个栈中，在函数执行后延迟执行，并且会拷贝一份变量。
	defer fmt.Println("num1: ", num1)
	defer fmt.Println("num2: ", num2)

	sum := num1 + num2

	num1++
	num2++
	fmt.Printf("sum: %d,num1: %d,num2: %d\n", sum, num1, num2)
	return sum
}
