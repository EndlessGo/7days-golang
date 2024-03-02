package geecache

// A ByteView holds an immutable view of bytes.
// ByteView 缓存值的抽象与封装
type ByteView struct {
	b []byte //选择byte类型是为了能够支持任意的数据类型的存储，例如字符串、图片等。
}

// Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	// 只读，拷贝一份
	return cloneBytes(v.b)
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
