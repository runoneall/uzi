package main

import (
	"fmt"
	"os"
	"uzi/client"
	"uzi/server"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("Usage: uzi <serve | connect> <addr> <auth>")
		return
	}

	mode := args[1]
	addr := args[2]
	auth := args[3]

	switch mode {

	case "serve":
		server.Serve(addr, auth)

	case "connect":
		client.Connect(addr, auth)

	default:
		fmt.Println("Not know mode:", mode)

	}
}
