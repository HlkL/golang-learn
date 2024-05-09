package main

import (
	"fmt"
	"time"
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
	selectDemo()
}

func selectDemo() {
	var ic chan int = make(chan int, 10)
	var sc chan string = make(chan string, 10)

	for i := 0; i < 10; i++ {
		ic <- i
		time.Sleep(time.Millisecond * 300)
	}

	for i := 0; i < 10; i++ {
		sc <- fmt.Sprintf("str %d", i)
		time.Sleep(time.Millisecond * 300)
	}

	// select: 执行任意一个符合条件的case，当没有符合条件则会发生 deadlock
	label:
	for {
		select {
		case num := <-ic:
			fmt.Println(num)
		case str := <-sc:
			fmt.Println(str)
		default:
			fmt.Println("管道没有数据")
			time.Sleep(time.Second)
			break label
		}
	}
}

func onlyReadOrWriteDemo() {
	var rc chan int = make(chan int, 2)
	var wc chan int = make(chan int, 2)
	onlyRead(rc)
	onlyWrite(wc)
}

func onlyRead(r <-chan int) {
	// 只读管道
	_ = <-r
	// r <- 12
}

func onlyWrite(w chan<- int) {
	// 只写管道
	w <- 32
	// _ = <-w
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
			v, ok := <-writer
			if !ok {
				break
			}
			fmt.Println("read data", v)
		}
		exit <- true
		close(exit)
	}(writeChan, exitChan)

	for {
		ok := <-exitChan
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
