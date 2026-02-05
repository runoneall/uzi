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

	return string(authPayload.MsgData) == *cli.Auth
}
