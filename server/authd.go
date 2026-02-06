package server

import (
	"net"
	"uzi/cli"
	"uzi/protocol"
)

func authd(conn net.Conn) bool {
	authPayload, err := protocol.Read(conn)
	if err != nil {
		return false
	}

	if authPayload.MsgType != "auth" {
		return false
	}

	ok := string(authPayload.MsgData) == *cli.Auth
	if !ok {
		protocol.Write(conn, protocol.Payload{
			MsgType: "auth",
			MsgData: []byte("failed"),
		})

	} else {
		protocol.Write(conn, protocol.Payload{
			MsgType: "auth",
			MsgData: []byte("success"),
		})
	}

	return ok
}
