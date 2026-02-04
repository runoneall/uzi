package server

import (
	"bufio"
	"encoding/json"
)

func getPayload(r *bufio.Reader) Payload {
	data, err := r.ReadBytes('\n')
	if err != nil {
		return Payload{Err: err}
	}

	var payload Payload
	if err := json.Unmarshal(data, &payload); err != nil {
		return Payload{Err: err}
	}

	return payload
}
