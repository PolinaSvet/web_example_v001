package mapEqualSave

import (
	"fmt"
	"sync"
	"time"
)

type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

type InMemoryCache struct {
	mu       sync.RWMutex
	expireIn time.Duration
	entries  map[string]CacheEntry
	stop     chan struct{} // Канал для остановки горутины удаления записей
}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		expireIn: expireIn,
		entries:  make(map[string]CacheEntry),
		stop:     make(chan struct{}),
	}
}

func (c *InMemoryCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = CacheEntry{
		settledAt: time.Now(),
		value:     value,
	}

	// Запускаем горутину для удаления записи по истечении времени
	go c.deleteEntry(key)
}

func (c *InMemoryCache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil
	}
	return entry.value
}

func (c *InMemoryCache) PrintCachedEntries() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	now := time.Now()
	for key, entry := range c.entries {
		remaining := c.expireIn - now.Sub(entry.settledAt)
		fmt.Printf("Key: %s, Value: %v, Время хранения: %s\n", key, entry.value, remaining.String())
	}
}

func (c *InMemoryCache) IsCachedEntries() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	res := false
	if len(c.entries) > 0 {
		res = true
	}

	return res
}

func (c *InMemoryCache) PrintCachedEntriesStr() string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	now := time.Now()
	outStr := ""
	for key, entry := range c.entries {
		remaining := c.expireIn - now.Sub(entry.settledAt)
		outStr += fmt.Sprintf("Key: %s, Value: %v Время хранения: %s;       ", key, entry.value, remaining.String())
	}
	return outStr
}

func (c *InMemoryCache) deleteEntry(key string) {
	for {
		select {
		case <-time.After(c.expireIn):
			c.mu.Lock()
			delete(c.entries, key)
			c.mu.Unlock()
			return
		case <-c.stop:
			return
		}
	}
}

func (c *InMemoryCache) StartDeletionCycle() {
	for {
		c.mu.RLock()
		hasEntries := len(c.entries) > 0
		c.mu.RUnlock()

		if !hasEntries {
			return
		}

		for key := range c.entries {
			go c.deleteEntry(key)
		}
		time.Sleep(c.expireIn)
	}
}

func (c *InMemoryCache) StopDeletionCycle() {
	close(c.stop)
}
