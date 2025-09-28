package main

import (
	"fmt"
	"net"
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
		resp := NewResp(connect)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(value)

		// テストのためrequestを無視し一時的にsOKを返す
		connect.Write([]byte("+OK\r\n"))
	}
}
