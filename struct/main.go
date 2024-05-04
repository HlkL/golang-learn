package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	name string
	age  int
}

type teacher struct {
	name string
	age  int
}

type Person struct {
	name string
	age  int
}

type Person1 struct {
	Name string
	Age  int
}

type person2 struct {
	Name string
	Age  int
}

type person3 struct {
	Name string
	age  int
}
// tag
type person4 struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}

func main() {

	// 声明学生结构体变量
	var s1 student
	s1.age = 13
	s1.name = "zhangsan"
	fmt.Println(s1)

	var s2 student = student{}
	s2.age = 14
	s2.name = "lisi"
	fmt.Println(s2)

	var s3 student = student{"wangwu",15}
	fmt.Println(s3)

	var s4 *student = new(student)
	(*s4).age = 16
	s4.name = "zhaoliu"
	fmt.Println(s4)

	var s5 *student = &student{}
	s5.age = 17
	s5.name = "chenqi"
	fmt.Println(s5)

	// 结构体序列化
	var teacher teacher = teacher{"tom",23}
	jsonStr, _ := json.Marshal(teacher)
	fmt.Println(string(jsonStr))

	var p *Person = new(Person)
	p.age = 13
	p.name = "zs"
	pjsonStr, _ := json.Marshal(p)
	fmt.Println(string(pjsonStr))

	var Person1 Person1 = Person1{"ls",24}
	p1jsonStr, _ := json.Marshal(Person1)
	fmt.Println(string(p1jsonStr))

	var person2 person2 = person2{"ww",25}
	p2jsonStr, _ := json.Marshal(person2)
	fmt.Println(string(p2jsonStr))

	var person3 person3 = person3{"zl",26}
	p3jsonStr, _ := json.Marshal(person3)
	fmt.Println(string(p3jsonStr))

	var person4 person4 = person4{"zl",26}
	p4jsonStr, _ := json.Marshal(person4)
	fmt.Println(string(p4jsonStr))

	// {zhangsan 13}
	// {lisi 14}
	// {wangwu 15}
	// &{zhaoliu 16}
	// &{chenqi 17}
	// {}
	// {}
	// {"Name":"ls","Age":24}
	// {"Name":"ww","Age":25}
	// {"Name":"zl"}
	// {"name":"zl","age":26}
}
