package client

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type serverMsg string

type model struct {
	conn     net.Conn
	auth     string
	viewport viewport.Model
	input    textinput.Model
	messages []string
}

func NewModel(conn net.Conn, auth string) model {
	ti := textinput.New()
	ti.Focus()

	return model{
		conn:     conn,
		auth:     auth,
		input:    ti,
		viewport: viewport.New(0, 0),
	}
}

func (m model) recvCmd() tea.Msg {
	str, err := bufio.NewReader(m.conn).ReadString('\n')
	if err != nil {
		return nil
	}

	return serverMsg(strings.TrimSpace(str))
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.recvCmd)
}

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

		return m, m.recvCmd

	}

	m.input, cmd = m.input.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf("--- 聊天室 ---\n%s\n\n%s", m.viewport.View(), m.input.View())
}
