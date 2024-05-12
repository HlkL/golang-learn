package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:24221")
	if err != nil {
		fmt.Println(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			break
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		if err == nil {
			conn.Write([]byte("hello!"))
			read := buf[:cnt]
			fmt.Printf("conn.RemoteAddr(): %v\n", conn.RemoteAddr())
			print(string(read[:cnt]))
		}
		
	}
}
