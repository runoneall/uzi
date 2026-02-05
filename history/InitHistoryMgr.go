package history

func InitHistoryMgr(size int) {
	once.Do(func() {
		Mgr = &HistoryMgr{
			data:    make([]string, size),
			maxSize: size,
		}
	})
}
