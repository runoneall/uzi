package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func newConn(conn net.Conn, auth string) {
	conn.(*net.TCPConn).SetNoDelay(true)
	w := bufio.NewWriter(conn)
	r := bufio.NewReader(io.LimitReader(conn, 1024*1024))

	defer func() {
		conn.Close()
		removeConn(conn)
	}()

	mu.Lock()
	Pool[conn] = struct{}{}
	mu.Unlock()

	remoteAddr, _ := net.ResolveTCPAddr("tcp", conn.RemoteAddr().String())
	remoteName := remoteAddr.IP.String()

	log := func(mode string, msg any) {
		fmt.Printf("[%s] %s: %v\n", remoteName, mode, msg)
	}

	mu.RLock()
	historySnapshot := make([]string, len(History))
	copy(historySnapshot, History)
	mu.RUnlock()

	count := 0
	for _, message := range historySnapshot {
		_, err := fmt.Fprintln(w, message)
		w.Flush()

		if err == nil {
			count++
		}
	}
	log("HISTORY", fmt.Sprintf("Sync %d messages", count))

	for {
		payload := getPayload(r)
		if payload.Err != nil {
			if payload.Err != io.EOF {
				log("ERROR", payload.Err)
			}

			return
		}

		if payload.Auth != auth {
			log("AUTH", "Invalid credentials")
			continue
		}

		message := strings.TrimSpace(payload.Message)
		log("BROADCAST", message)
		broadcast(fmt.Sprintf("[%s] %s", remoteName, message))
	}
}
