package client

import (
	"uzi/protocol"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) recvnext() tea.Cmd {
	return func() tea.Msg {
		msgPayload, err := protocol.Read(m.conn)
		if err != nil {
			return nil
		}

		if msgPayload.MsgType != "message" {
			return nil
		}

		return msgPayload
	}
}
