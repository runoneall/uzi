package client

import (
	"log/slog"
	"net"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func startui(conn net.Conn) {
	ti := textinput.New()
	ti.Placeholder = "输入消息..."
	ti.Focus()

	m := &model{
		conn:     conn,
		viewport: viewport.New(0, 0),
		input:    ti,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		slog.Error(err.Error())
	}
}
