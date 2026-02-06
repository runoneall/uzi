package server

import (
	"log/slog"
	"net"
	uziconn "uzi/conn"
	"uzi/history"
	"uzi/protocol"
)

func onconn(conn net.Conn) {
	defer conn.Close()

	if !authd(conn) {
		slog.Error("authenticate failed", "addr", conn.RemoteAddr())
		return
	}

	slog.Info("new client joined", "addr", conn.RemoteAddr())

	for _, historyMsg := range history.Mgr.Get() {
		protocol.Write(conn, protocol.Payload{
			MsgType: "message",
			MsgData: []byte(historyMsg),
		})
	}
	uziconn.Mgr.Add(conn)
	msgloop(conn)
}
