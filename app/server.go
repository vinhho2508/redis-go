package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	// "time"
	// "bytes"
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
	for {
		_, err = c.Read(buf)
		if err != nil {
			break
		}
	
		log.Printf("Read: %s", buf)
		_, err = c.Write([]byte("+PONG\r\n"))
		// if bytes.Equal(buf, []byte("*1\r\n$4\r\nPING\r\n")) {
		// 	if err != nil {
		// 		fmt.Println("Can not write")
		// 	}
		// } else {
		// 	_,err = c.Write([]byte("+NO\r\n"))
		// 	if err != nil {
		// 		fmt.Println("Can not write")
		// 	}
		// }

		// time.Sleep(time.Minute)

		}
}
