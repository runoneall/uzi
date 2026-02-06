package client

import (
	"uzi/protocol"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		return m.keyevent(msg)

	case tea.WindowSizeMsg:
		m.resizewin(msg)

	case protocol.Payload:
		return m.onrecv(msg)

	}

	return m, nil
}
