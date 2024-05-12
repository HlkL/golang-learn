package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:24221")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer conn.Close()

	read := bufio.NewReader(os.Stdin)
	str, err := read.ReadString('\n')
	if err == nil {
		conn.Write([]byte(str))
	} else {
		fmt.Printf("err: %v\n", err)
	}
}
