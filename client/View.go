package client

import "fmt"

func (m *model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.input.View(),
	)
}
