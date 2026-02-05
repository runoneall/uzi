package client

import (
	"bufio"
	"net"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
)

func newModel(conn net.Conn, auth string) model {
	ti := textinput.New()
	ti.Focus()

	return model{
		conn:     conn,
		reader:   bufio.NewReader(conn),
		auth:     auth,
		input:    ti,
		viewport: viewport.New(0, 0),
	}
}
