package history

func (hm *HistoryMgr) Get() []string {
	hm.mu.RLock()
	defer hm.mu.RUnlock()

	if !hm.isFull {
		return append([]string(nil), hm.data[:hm.head]...)
	}

	data := make([]string, hm.maxSize)
	copy(data, hm.data[hm.head:])
	copy(data[hm.maxSize-hm.head:], hm.data[:hm.head])
	return data
}
