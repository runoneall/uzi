package client

import (
	"encoding/json"
	"fmt"
	"net"
	"uzi/server"
)

func send(conn net.Conn, auth, message string) {
	data, _ := json.Marshal(server.Payload{Auth: auth, Message: message})
	fmt.Fprintln(conn, string(data))
}
