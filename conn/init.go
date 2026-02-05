package conn

import (
	"net"
	"sync"
)

type ConnItem struct {
	Conn net.Conn
	ID   int64
}

type ConnMgr struct {
	mu        sync.RWMutex
	conns     []*ConnItem
	m         map[int64]int
	idCounter int64
}

var Mgr = &ConnMgr{
	conns: make([]*ConnItem, 0, 100),
	m:     make(map[int64]int),
}
