package geecache

import pb "geecache/geecachepb"

// PeerPicker 节点选择接口
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool) //根据传入的key选择相应节点
}

// PeerGetter 节点获取（缓存值）接口
type PeerGetter interface {
	// 将http通信的中间载体换成protobuf
	Get(in *pb.Request, out *pb.Response) error //从对应group查找缓存值
}
