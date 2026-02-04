package server

import "net"

func removeConn(conn net.Conn) {
	mu.Lock()
	delete(Pool, conn)
	mu.Unlock()
}
