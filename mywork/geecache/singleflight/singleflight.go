package singleflight

import "sync"

// call is an in-flight or completed Do call
// call是一个正在进行中或者已经完成的请求
type call struct {
	wg  sync.WaitGroup // 锁避免重入
	val interface{}
	err error
}

// Group 管理不同key的请求(call)
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

// Do 保证对于相同的key，无论Do被调用多少次，fn都只会被调用一次
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	// 如果请求正在进行中，则等待
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	// 如果请求还没有进行，则新建一个请求
	c := new(call)
	// 发起请求前，加锁
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()
	return c.val, c.err
}
