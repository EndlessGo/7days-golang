// 基于HTTP实现分布式缓存需要的节点间通信

package geecache

import (
	"fmt"
	"geecache/consistenthash"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var _ PeerGetter = (*httpGetter)(nil)

type httpGetter struct {
	baseURL string // 远程节点地址，如 http://example.com/_geecache/
}

// Get 通过http获取缓存值
func (h *httpGetter) Get(group string, key string) ([]byte, error) {
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		url.QueryEscape(group),
		url.QueryEscape(key))
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.Status)
	}
	//bytes, err := ioutil.ReadAll(res.Body)
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}
	return bytes, nil
}

const (
	defaultPath     = "/_geecache/"
	defaultReplicas = 50
)

type HTTPPool struct {
	self        string     // 节点url，https://域名或ip:端口
	basePath    string     // 主机可能承载其他服务，加一个节点通信地址前缀
	mu          sync.Mutex // guards peers and httpGetters
	peers       *consistenthash.Map
	httpGetters map[string]*httpGetter // key是远程节点的URL
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultPath,
	}
}

func (p *HTTPPool) Log(format string, v ...interface{}) {
	// 记得展开v...，而不是v，否则%s %s 会少匹配第二个，多输出%!s(MISSING)
	// https://blog.csdn.net/molaifeng/article/details/126497350
	// 解决 golang 中输出 (MISSING)
	log.Printf("[Server %s] %s\n", p.self, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		http.Error(w, "HTTPPool serving unexpected path: "+r.URL.Path, http.StatusNotFound)
	}
	p.Log("%s %s", r.Method, r.URL.Path)
	// 相对路径格式 /basePath/groupName/key
	// 前缀后的部分groupName/key，按"/"分割成2个子串
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupName := parts[0]
	key := parts[1]
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusBadRequest)
		return
	}
	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = w.Write(view.ByteSlice())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = consistenthash.New(defaultReplicas, nil)
	p.peers.Add(peers...)
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}
func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}
