package geecache

// PeerPicker 节点选择接口
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool) //根据传入的key选择相应节点
}

// PeerGetter 节点获取（缓存值）接口
type PeerGetter interface {
	Get(group string, key string) ([]byte, error) //从对应group查找缓存值
}
