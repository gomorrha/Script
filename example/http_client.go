package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("error dialing", err.Error())
		return
	}

	defer conn.Close()

	msg := "GET/HTTP/1.1\r\n"
	msg += "Host: www.baidu.com\r\n"
	msg += "Connection: close\r\n"
	msg += "\r\n\r\n"

	_, err = io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed, ", err)
		return
	}

	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		if err != nil || count == 0 {
			break
		}
		fmt.Println(string(buf[0:count]))
	}
}
