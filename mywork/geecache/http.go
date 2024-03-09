// 基于HTTP实现分布式缓存需要的节点间通信

package geecache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const defaultPath = "/_geecache/"

type HTTPPool struct {
	self      string // 节点url，https://域名或ip:端口
	cachePath string // 主机可能承载其他服务，加一个节点通信地址前缀
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:      self,
		cachePath: defaultPath,
	}
}

func (p *HTTPPool) Log(format string, v ...interface{}) {
	// 记得展开v...，而不是v，否则%s %s 会少匹配第二个，多输出%!s(MISSING)
	log.Printf("[Server %s] %s\n", p.self, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, p.cachePath) {
		http.Error(w, "HTTPPool serving unexpected path: "+r.URL.Path, http.StatusNotFound)
	}
	p.Log("%s %s", r.Method, r.URL.Path)
	// 相对路径格式 /cachePath/groupName/key
	// 前缀后的部分groupName/key，按"/"分割成2个子串
	parts := strings.SplitN(r.URL.Path[len(p.cachePath):], "/", 2)
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
