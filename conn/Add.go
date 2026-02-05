package conn

import (
	"net"
	"sync/atomic"
)

func (cm *ConnMgr) Add(conn net.Conn) int64 {
	newID := atomic.AddInt64(&cm.idCounter, 1)

	cm.mu.Lock()
	defer cm.mu.Unlock()

	newItem := &ConnItem{
		Conn: conn,
		ID:   newID,
	}

	cm.m[newID] = len(cm.conns)
	cm.conns = append(cm.conns, newItem)

	return newID
}
