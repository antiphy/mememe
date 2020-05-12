package models

import (
	"sync"
)

type QueryParams struct {
	ID       int
	Page     int
	PageSize int
	Total    int
}

type Cache interface {
	GET(key string) *Setting
	SET(key string, val *Setting)
}

type cache struct {
	data map[string]*Setting
	*sync.RWMutex
}

func (c *cache) GET(key string) *Setting {
	c.RLock()
	val, exists := c.data[key]
	c.RUnlock()
	if !exists {
		return &Setting{
			SettingKey:   key,
			SettingValue: "default setting value, please set this field.",
		}
	}
	return val
}

func (c *cache) SET(key string, val *Setting) {
	c.Lock()
	c.data[key] = val
	c.Unlock()
}

func NewCache() Cache {
	return &cache{
		make(map[string]*Setting),
		new(sync.RWMutex),
	}
}
