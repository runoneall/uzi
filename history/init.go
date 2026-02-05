package history

import "sync"

type HistoryMgr struct {
	mu      sync.RWMutex
	data    []string
	maxSize int
	head    int
	isFull  bool
}

var (
	once sync.Once
	Mgr  *HistoryMgr
)
