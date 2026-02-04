package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"uzi/server"
)

func sendui(conn net.Conn, auth string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		if text == "exit" {
			break
		}

		payload := server.Payload{
			Auth:    auth,
			Message: text,
		}

		data, err := json.Marshal(payload)
		if err != nil {
			panic(err)
		}

		_, err = fmt.Fprintln(conn, string(data))
		if err != nil {
			fmt.Println("! 网络问题，发送失败")
			break
		}
	}
}
