package client

import (
	"bufio"
	"net"
	"strings"
)

func recv(conn net.Conn, out chan string) {
	r := bufio.NewReader(conn)
	for {
		message, err := r.ReadString('\n')
		if err != nil {
			return
		}

		out <- strings.TrimSpace(message)
	}
}
