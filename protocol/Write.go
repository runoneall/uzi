package protocol

import (
	"encoding/binary"
	"io"
)

func Write(w io.Writer, data Payload) error {
	typeBytes := []byte(data.MsgType)
	if err := binary.Write(w, binary.BigEndian, uint64(len(typeBytes))); err != nil {
		return err
	}

	if _, err := w.Write(typeBytes); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint64(len(data.MsgData))); err != nil {
		return err
	}

	if _, err := w.Write(data.MsgData); err != nil {
		return err
	}

	return nil
}
