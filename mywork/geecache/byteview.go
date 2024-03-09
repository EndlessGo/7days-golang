// ByteView 对缓存值的抽象，返回拷贝是为了防止缓存值被外部程序修改

package geecache

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice 返回一个拷贝，防止缓存值被外部程序修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
