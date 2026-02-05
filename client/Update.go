package client

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "esc":
			return m, tea.Quit

		case "enter":
			if v := m.input.Value(); v != "" {
				send(m.conn, m.auth, v)
				m.input.Reset()
			}

		}

	case tea.WindowSizeMsg:
		headerHeight := 2
		inputHeight := 3
		verticalMarginHeight := headerHeight + inputHeight

		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - verticalMarginHeight

		m.input.Width = msg.Width - 5

	case serverMsg:
		m.messages = append(m.messages, string(msg))

		m.viewport.SetContent(strings.Join(m.messages, "\n"))
		m.viewport.GotoBottom()

		return m, m.recv

	}

	m.input, cmd = m.input.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
