package client

import (
	"fmt"
	"net"
)

func Connect(addr string, auth string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("已连接到服务器。输入消息并回车发送 (输入 'exit' 退出)")
	go recv(conn)
	sendui(conn, auth)
}
