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

const MaxHistory = 100

var mu = sync.RWMutex{}

var Pool = make(map[net.Conn]struct{})
var History []string
