package client

import (
	"bufio"
	"fmt"
	"net"
)

func recv(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		message, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Print(message)
	}
}
