package main

import (
	"fmt"
	_ "time"
)

func print() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello,world")

	}
}

var intChanel = make(chan int, 200)

func init() {
	for i := 0; i < 200; i++ {
		intChanel <- i * 2
	}
}

func main() {
	readAndWriteDemo()
}

func readAndWriteDemo() {
	var writeChan chan int = make(chan int, 50)
	var exitChan chan bool = make(chan bool, 1)

	go func(writer chan int) {
		for i := 0; i < 50; i++ {
			writeChan <- i
			fmt.Println("write data", i)
		}
		close(writer)
	}(writeChan)

	go func(writer chan int, exit chan bool) {
		for {
			v, ok := <- writer
			if !ok {
				break
			}
			fmt.Println("read data", v)
		}
		exit <- true
		close(exit)
	}(writeChan, exitChan)

	for {
		ok := <- exitChan
		if !ok {
			break
		}
	}
}

func channel() {
	// 管道定义
	var intChanel = make(chan int, 3)

	intChanel <- 123
	intChanel <- 1243
	intChanel <- 1243

	num := <-intChanel
	fmt.Println(num)
}

func traversal() {
	close(intChanel)
	// 在遍历前如果管道没有关闭，会出现死锁
	for num := range intChanel {
		fmt.Println(num)
	}
}
