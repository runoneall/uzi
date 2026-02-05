package server

import (
	"fmt"
	"net"
)

func Serve(addr string, auth string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("TCP Serve at:", addr)
	fmt.Printf("Use 'uzi connect %s %s' to connect this server\n", listener.Addr().String(), auth)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept failed:", err)
			continue
		}

		go newConn(conn, auth)
	}
}
