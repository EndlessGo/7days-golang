package lru

import "container/list"

// Cache is a LRU cache. It is not safe for concurrent access.
type Cache struct {
	// go标准库实现的双向链表
	ll *list.List
	// 字典，键是字符串，而不是
	cache map[string]*list.Element
	// 感觉不需要 maxBytes和nBytes两个字段，直接一个size即可
	// 最大字节数
	maxBytes int64
	// 当前已使用的内存
	nBytes int64
	// optional and executed when an entry is purged.
	// 记录被移除时的回调
	OnEvicted func(key string, value Value)
}

// 链表节点元素
type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

// New is the Constructor of Cache
// 实例化
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Add adds a value to the cache.
// 插入：包括新增或者是更新
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		// 更新
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		// 新增
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		// 移除
		c.RemoveOldest()
	}
}

// Get look ups a key's value
// 查找，更新节点到队头
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item
// 删除队尾，移除最近最少访问的节点
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Len the number of cache entries
func (c *Cache) Len() int {
	return c.ll.Len()
}
