package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func newConn(conn net.Conn, auth string) {
	defer func() {
		conn.Close()
		removeConn(conn)
	}()

	mu.Lock()
	Pool[conn] = struct{}{}
	mu.Unlock()

	log := func(mode string, msg any) {
		fmt.Printf("[%s] %s: %v\n", conn.RemoteAddr(), mode, msg)
	}

	limitedReader := io.LimitReader(conn, 1024*1024)
	r := bufio.NewReader(limitedReader)
	for {
		payload := getPayload(r)
		if payload.Err != nil {
			if payload.Err != io.EOF {
				log("ERROR", payload.Err)
			}

			return
		}

		if payload.Auth != auth {
			log(" AUTH", "Invalid credentials")
			continue
		}

		message := strings.TrimSpace(payload.Message)
		log("BROAD", message)
		broadcast(fmt.Sprintf("[%s] %s", conn.RemoteAddr(), message))
	}
}
