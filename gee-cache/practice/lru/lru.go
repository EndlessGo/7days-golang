package lru

import "container/list"

// element 链表节点元素
type element struct {
	key   string     // 缓存的键
	value CacheValue // 缓存的值
}

// CacheValue 缓存值的抽象接口
type CacheValue interface {
	Len() int
}

// Cache lru缓存,非并发安全
type Cache struct {
	maxBytes  int64                    // 最大内存，用字节而不是链表长度，能更准确的限制大小
	bytes     int64                    // 当前内存
	ll        *list.List               // 双向链表
	cache     map[string]*list.Element // 缓存，k: string，v:链表节点指针
	onEvicted func(key string, value CacheValue)
}

// New 实例化
func New(maxBytes int64, onEvicted func(string, CacheValue)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		onEvicted: onEvicted,
	}
}

// Get 获取缓存值
func (c *Cache) Get(key string) (value CacheValue, ok bool) {
	if ele, ok := c.cache[key]; ok {
		// 更新元素到队首，队尾是最近最少访问队数据
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*element)
		return kv.value, true
	}
	return
}

// RemoveOldest 删除最近最少访问的节点
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		// 删除队尾元素
		c.ll.Remove(ele)
		kv := ele.Value.(*element)
		delete(c.cache, kv.key)
		// 更新大小
		c.bytes -= int64(len(kv.key) + kv.value.Len())
		if c.onEvicted != nil {
			c.onEvicted(kv.key, kv.value)
		}
	}
}

// Add 新增/修改缓存值
func (c *Cache) Add(key string, value CacheValue) {
	if ele, ok := c.cache[key]; ok {
		// 更新队首
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*element)
		// 更新大小
		c.bytes += int64(value.Len() - kv.value.Len())
		// 更新缓存值
		kv.value = value
	} else {
		// 插入队首
		ele = c.ll.PushFront(&element{key: key, value: value})
		c.bytes += int64(len(key) + value.Len())
		c.cache[key] = ele
	}
	// 统一删除队尾元素
	for c.maxBytes != 0 && c.maxBytes < c.bytes {
		c.RemoveOldest()
	}
}

// Len 缓存节点数
func (c *Cache) Len() int {
	return c.ll.Len()
}
