package server

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	uziconn "uzi/conn"
	"uzi/history"
	"uzi/protocol"
)

func msgloop(conn net.Conn) {
	msgSender, _ := net.ResolveTCPAddr("tcp", conn.RemoteAddr().String())
	msgSenderIP := msgSender.IP.String()

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

		msgString := fmt.Sprintf("[%s] %s", msgSenderIP, string(msgPayload.MsgData))
		history.Mgr.Add(msgString)

		slog.Info("broadcast", "message", msgString)
		uziconn.Mgr.Broadcast(func(conn net.Conn) bool {
			err := protocol.Write(conn, protocol.Payload{
				MsgType: "message",
				MsgData: []byte(msgString),
			})

			return err == nil
		})
	}
}
