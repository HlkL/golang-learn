package main

import (
	"errors"
	"fmt"
)

func main() {
	throwError()
}

func catchError() int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	num1 := 1
	num2 := 0
	return num1 / num2
}

func customizeError(name string) error {
	if name == "" {
		return errors.New("name is null")
	}
	return nil
}

func throwError()  {
	err := customizeError("")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}