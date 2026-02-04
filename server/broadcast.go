package server

import (
	"fmt"
	"net"
	"time"
)

func broadcast(message string) {
	mu.RLock()
	conns := make([]net.Conn, 0, len(Pool))
	for conn := range Pool {
		conns = append(conns, conn)
	}
	mu.RUnlock()

	for _, conn := range conns {
		go func(conn net.Conn) {
			conn.SetWriteDeadline(time.Now().Add(5 * time.Second))

			_, err := fmt.Fprintln(conn, message)
			if err != nil {
				conn.Close()
				removeConn(conn)
			}
		}(conn)
	}
}
