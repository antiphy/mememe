package models

import "sync"

type blocker struct {
	data map[string]int
	*sync.RWMutex
}

type Blocker interface {
	Incr(key string)
	IsBlocked(key string) bool
}

func NewBlocker() Blocker {
	return &blocker{
		make(map[string]int),
		new(sync.RWMutex),
	}
}

func (b *blocker) Incr(key string) {
	b.Lock()
	count, exists := b.data[key]
	if !exists {
		b.data[key] = 1
	} else {
		b.data[key] = count + 1
	}
	b.Unlock()
}

func (b *blocker) IsBlocked(key string) bool {
	b.RLock()
	count, exists := b.data[key]
	b.RUnlock()
	if !exists {
		return false
	}
	if count >= 3 {
		return true
	}
	return false
}
