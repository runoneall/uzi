package client

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) recv() tea.Msg {
	str, err := m.reader.ReadString('\n')
	if err != nil {
		return nil
	}

	return serverMsg(strings.TrimSpace(str))
}
