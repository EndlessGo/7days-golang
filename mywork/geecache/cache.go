// cache实现了一个互斥读写的并发缓存

package geecache

import (
	"geecache/lru"
	"sync"
)

type cache struct {
	mu       sync.RWMutex //读写锁替换了互斥锁sync.Mutex
	lru      *lru.Cache
	maxBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.maxBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
