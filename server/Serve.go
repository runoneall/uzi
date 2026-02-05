package server

import (
	"log/slog"
	"net"
	"uzi/cli"
	"uzi/history"
)

func Serve() {
	history.InitHistoryMgr(1000)

	address := net.JoinHostPort(*cli.Host, *cli.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	slog.Info("server started", "addr", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		go onconn(conn)
	}
}
