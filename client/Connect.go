package client

import (
	"net"

	tea "github.com/charmbracelet/bubbletea"
)

func Connect(addr string, auth string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	p := tea.NewProgram(newModel(conn, auth))
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
