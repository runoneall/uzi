package server

import (
	"net"
	"sync"
)

type Payload struct {
	Auth    string `json:"auth"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

var mu = sync.RWMutex{}
var Pool = make(map[net.Conn]struct{})
