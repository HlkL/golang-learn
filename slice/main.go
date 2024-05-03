package main

import (
	"fmt"
)

// 切片是引用类型，数组是值类型。切片内存空间可以动态的变化
func main() {
	// 切片定义: 数组引用
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	// 左闭右开
	arrRef := arr[0:3]
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("arrRef: %v\n", arrRef)
	// 修改数组的元素
	modifyArrItem(arr)
	fmt.Printf("after arr modify: %v\n", arr)
	// 传入数组指针修改数组元素
	modifyByPointer(&arr)
	fmt.Printf("use pointer modify after arr: %v\n", arr)

	// 修改切片，会改变原数组值
	arrRef[0] = 45
	fmt.Printf("modify origin arr after slice: %v\n", arr)

	// 修改原数组的值，切片也会跟着改变
	arr[0] = 34
	fmt.Printf("arrRef: %v\n", arrRef)

	// 使用make定义切片
	var slice []int = make([]int, 5, 10)
	fmt.Printf("slice: %v\n", slice)

	// 方式三
	var slice2 []int = []int{1,2,3,4,5}
	fmt.Printf("slice2: %v\n", slice2)
	
	// 执行结果
	// arr: [1 2 3 4 5]
	// arrRef: [1 2 3]
	// after arr modify: [1 2 3 4 5]
	// use pointer modify after arr: [123 2 3 4 5]
	// modify origin arr after slice: [45 2 3 4 5]
	// arrRef: [34 2 3]
	// slice: [0 0 0 0 0]
	// slice2: [1 2 3 4 5]
}

// 函数形参接收数组必须写明数组长度
func modifyArrItem(arr [5]int) {
	arr[0] = 123
}

func modifyByPointer(arr *[5]int) {
	(*arr)[0] = 123
}
