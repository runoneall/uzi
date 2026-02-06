package client

import (
	"log/slog"
	"net"
	"uzi/cli"
	"uzi/protocol"
)

func authconn(conn net.Conn) bool {
	protocol.Write(conn, protocol.Payload{
		MsgType: "auth",
		MsgData: []byte(*cli.Auth),
	})

	authPayload, err := protocol.Read(conn)
	if err != nil {
		slog.Error(err.Error())
		return false
	}

	return authPayload.MsgType == "auth" && string(authPayload.MsgData) == "success"
}
