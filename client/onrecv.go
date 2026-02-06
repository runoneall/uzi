package client

import (
	"strings"
	"uzi/protocol"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) onrecv(msg protocol.Payload) (tea.Model, tea.Cmd) {
	m.messages = append(m.messages, string(msg.MsgData))
	m.viewport.SetContent(strings.Join(m.messages, "\n"))
	m.viewport.GotoBottom()

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)

	return m, tea.Batch(cmd, m.recvnext())
}
