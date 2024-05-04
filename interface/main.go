package main

import "fmt"

// go 语言实现接口没有 implement 关键字，只需要定义与接口名称一致，且定义与所实现接口的全部抽象方法即可。鸭子类型
type game interface {
	start()
	end()
}

type lol struct {

}

func (lol lol) start() {
	fmt.Println("开始游戏")
}

func (lol lol) end() {
	fmt.Println("结束游戏")
}

type computer struct{

}

func (computer *computer) work(game game) {
	game.start()
	game.end()
}

func work(game game)  {
	game.start()
	game.end()
}

func main()  {
	lol := lol{}
	computer := computer{}
	computer.work(lol)
	work(lol)
}