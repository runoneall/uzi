package client

import (
	"log/slog"
	"net"
	"uzi/cli"
)

func Connect() {
	address := net.JoinHostPort(*cli.Host, *cli.Port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	if !authconn(conn) {
		slog.Error("authenticate failed", "addr", conn.RemoteAddr(), "auth", *cli.Auth)
		return
	}

	startui(conn)
}
