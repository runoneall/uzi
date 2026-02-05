package history

func (hm *HistoryMgr) Add(entry string) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	hm.data[hm.head] = entry
	hm.head = (hm.head + 1) % hm.maxSize

	if hm.head == 0 {
		hm.isFull = true
	}
}
