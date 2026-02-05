package protocol

import (
	"encoding/binary"
	"io"
)

func Read(r io.Reader) (Payload, error) {
	var data Payload

	var typeLen uint64
	if err := binary.Read(r, binary.BigEndian, &typeLen); err != nil {
		return data, err
	}

	typeBuf := make([]byte, typeLen)
	if _, err := io.ReadFull(r, typeBuf); err != nil {
		return data, err
	}
	data.MsgType = string(typeBuf)

	var dataLen uint64
	if err := binary.Read(r, binary.BigEndian, &dataLen); err != nil {
		return data, err
	}

	data.MsgData = make([]byte, dataLen)
	if _, err := io.ReadFull(r, data.MsgData); err != nil {
		return data, err
	}

	return data, nil
}
