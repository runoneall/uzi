package conn

import (
	"net"
	"sync"
)

func (cm *ConnMgr) Broadcast(f func(conn net.Conn) bool) {
	cm.mu.RLock()
	items := make([]*ConnItem, len(cm.conns))
	copy(items, cm.conns)
	cm.mu.RUnlock()

	if len(items) == 0 {
		return
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, 100)
	failedIDs := make(chan int64, len(items))

	for _, item := range items {
		wg.Add(1)
		sem <- struct{}{}

		go func(item *ConnItem) {
			defer wg.Done()
			defer func() { <-sem }()

			if !f(item.Conn) {
				failedIDs <- item.ID
			}
		}(item)
	}

	wg.Wait()
	close(failedIDs)

	for id := range failedIDs {
		cm.Remove(id)
	}
}
