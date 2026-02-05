package client

import (
	"fmt"
)

func (m model) View() string {
	return fmt.Sprintf("--- 聊天室 ---\n%s\n\n%s", m.viewport.View(), m.input.View())
}
