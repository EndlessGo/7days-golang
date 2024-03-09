package geecache

import (
	"fmt"
	"log"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

// GetterFunc 接口型函数，用于用户自定义从不同的数据源（文件、数据库等）获取数据并添加到函数
type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// Group 组缓存，即具名缓存空间，管理并发lru缓存
// 是gee cache最核心的数据结构，负责与外部交互，缓存存储和获取等流程
type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

// Get 从缓存中获取值
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	if v, ok := g.mainCache.get(key); ok {
		log.Printf("[GeeCache] hit, key: %s, value: %s", key, v.String())
		return v, nil
	}
	return g.load(key)
}

func (g *Group) load(key string) (ByteView, error) {
	// TODO: 是否从远程节点获取
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
