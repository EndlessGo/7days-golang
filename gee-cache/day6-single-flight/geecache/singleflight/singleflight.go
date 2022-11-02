package singleflight

import "sync"

// call is an in-flight or completed Do call
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		// 如果请求正在进行中，则等待
		// 请求结束，返回结果
		c.wg.Wait()
		return c.val, c.err
	}
	c := new(call)
	// 发起请求前加锁
	c.wg.Add(1)
	// 发起请求前加锁
	g.m[key] = c
	g.mu.Unlock()

	// 调用 fn，发起请求
	c.val, c.err = fn()
	// 请求结束
	c.wg.Done()

	g.mu.Lock()
	// 更新g.m
	delete(g.m, key)
	g.mu.Unlock()

	// 返回结果
	return c.val, c.err
}
