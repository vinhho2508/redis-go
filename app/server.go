package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func stop_server(l io.Closer) {
	err := l.Close()
	if err != nil {
		fmt.Println("close server error")
	}
}
func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	c, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer stop_server(c)

	buf := make([]byte, 128)
	_, err = c.Read(buf)
	if err != nil {
		fmt.Println("Can not read")
	}

	log.Printf("Read: %s", buf)
	_, err = c.Write([]byte("+PONG\r\n"))
	if err != nil {
		fmt.Println("Can not write")
	}
}
