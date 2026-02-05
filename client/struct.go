package client

import (
	"bufio"
	"net"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
)

type serverMsg string

type model struct {
	conn     net.Conn
	reader   *bufio.Reader
	auth     string
	viewport viewport.Model
	input    textinput.Model
	messages []string
}
