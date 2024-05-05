package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	readFile()
	
}

func readFile() {
	file, _ := os.Open("./main.go")
	defer file.Close()

	buf := bufio.NewReader(file)

	for {
		str, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	fmt.Println("===================================================")

	byte, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(string(byte))
	fmt.Println("===================================================")

	rst, err := os.ReadFile("./main.go")
	if err != nil {
		return 
	}
	fmt.Print(string(rst))
}