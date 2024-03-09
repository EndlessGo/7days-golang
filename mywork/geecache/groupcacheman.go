package geecache

import "sync"

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group) // 全局变量，用来存储所有的Group
)

// NewGroup 创建一个新的Group实例
func NewGroup(name string, maxBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:   name,
		getter: getter,
		mainCache: cache{
			maxBytes: maxBytes,
		},
	}
	groups[name] = g
	return g
}

// GetGroup 获取指定名称的Group
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}
