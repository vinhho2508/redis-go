package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func stopServer(l io.Closer) {
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

	defer stopServer(l)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 128)
	for {
		n, err := c.Read(buf)
		if err != nil {
			return
		}

		log.Printf("Read: %s", buf[:n])
		if _, err := c.Write([]byte("+PONG\r\n")); err != nil {
			fmt.Println("Error write output: ", err.Error())
		}
	}
}
