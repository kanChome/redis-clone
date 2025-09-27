package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Listening on port :6379")

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	connect, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connect.Close()

	for {
		buffer := make([]byte, 1024)

		_, err = connect.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		connect.Write([]byte("+OK\r\n"))
	}
}
