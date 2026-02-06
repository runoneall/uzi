package client

import "uzi/protocol"

func (m *model) onsend() {
	if v := m.input.Value(); v != "" {
		protocol.Write(m.conn, protocol.Payload{
			MsgType: "message",
			MsgData: []byte(v),
		})

		m.input.Reset()
	}
}
