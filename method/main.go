package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (person Person) sayHello() {
	fmt.Println("hello,", person.Name)
}

func (person *Person) setName(name string) {
	person.Name = name
}

func (person Person) String() string {
	return fmt.Sprintf("person = {Name : %s, Age : %d}", person.Name, person.Age)
}
func main() {
	person := Person{"zhansan", 12}
	person.sayHello()
	fmt.Println(person)
	person.setName("lisi")
	fmt.Println(person)
	fmt.Println(&person)
}
