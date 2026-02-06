package client

import (
	"net"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	conn     net.Conn
	messages []string
	viewport viewport.Model
	input    textinput.Model
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		m.recvnext(),
	)
}
