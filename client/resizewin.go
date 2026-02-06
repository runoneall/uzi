package client

import tea "github.com/charmbracelet/bubbletea"

func (m *model) resizewin(msg tea.WindowSizeMsg) {
	m.viewport.Width = msg.Width

	inputHeight := 3
	m.viewport.Height = msg.Height - inputHeight
}
