package consistenthash

import (
	"hash/crc32"
	"slices"
	"strconv"
)

// 一致性哈希算法（环状哈希），解决缓存雪崩问题
// 虚拟节点，解决当服务器节点过少时的数据倾斜问题

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int            // 虚拟节点倍数
	keys     []int          // 一致性哈希环，有序，元素是虚拟节点的哈希值
	hashMap  map[int]string // 虚拟节点与真实节点的映射表, k:虚拟节点的哈希值, v:真实节点的名称
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add 添加缓存节点，keys 是节点名称，会对每个节点创建 replicas 个虚拟节点。
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	// sort.Ints(m.keys)
	// faster
	slices.Sort(m.keys)
}

// Get 根据数据获取最近的缓存节点名称（真实节点）。
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	//sort.Search(len(m.keys), func(i int) bool {
	//	return m.keys[i] >= hash
	//})
	idx, _ := slices.BinarySearch(m.keys, hash)
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
