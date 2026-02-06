package client

import tea "github.com/charmbracelet/bubbletea"

func (m *model) keyevent(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {

	case "ctrl+c", "esc":
		return m, tea.Quit

	case "enter":
		m.onsend()

	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
