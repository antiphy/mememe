package models

import (
	"errors"
	"sync"
	"time"
)

type QueryParams struct {
	ID       int
	Page     int
	PageSize int
	Total    int
}

type JWTClaims struct {
	UID      int
	Name     string
	ExpireTS int64
	Issuer   string
}

func (c JWTClaims) Valid() error {
	if c.ExpireTS <= time.Now().Unix() {
		return errors.New("token expired")
	}
	return nil
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
