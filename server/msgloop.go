package server

import (
	"io"
	"log/slog"
	"net"
	uziconn "uzi/conn"
	"uzi/history"
	"uzi/protocol"
)

func msgloop(conn net.Conn) {
	for {
		msgPayload, err := protocol.Read(conn)
		if err != nil {
			if err == io.EOF {
				slog.Info("connection closed by remote", "addr", conn.RemoteAddr())

			} else {
				slog.Error(
					"read message failed",
					"addr", conn.RemoteAddr(),
					"err", err,
				)
			}

			return
		}

		if msgPayload.MsgType != "message" {
			slog.Warn("invalid message type", "addr", conn.RemoteAddr())
			continue
		}

		history.Mgr.Add(string(msgPayload.MsgData))
		uziconn.Mgr.Broadcast(func(conn net.Conn) bool {
			_, err := conn.Write(msgPayload.MsgData)
			return err == nil
		})
	}
}
