package conn

func (cm *ConnMgr) Remove(id int64) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	idx, ok := cm.m[id]
	if !ok {
		return
	}

	lastIdx := len(cm.conns) - 1
	if idx != lastIdx {
		lastItem := cm.conns[lastIdx]
		cm.conns[idx] = lastItem
		cm.m[lastItem.ID] = idx
	}

	cm.conns[lastIdx] = nil
	cm.conns = cm.conns[:lastIdx]
	delete(cm.m, id)
}
