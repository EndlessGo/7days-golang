package codec

import (
	"io"
)

type Header struct {
	// 服务名+函数名
	ServiceMethod string // format "Service.Method"
	// 请求的序列
	Seq uint64 // sequence number chosen by client
	// 错误
	Error string
}

// Codec 对消息体进行编解码的接口
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

// NewCodecFunc 构造函数
type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

// NewCodecFuncMap 客户端和服务器可以通过Codec的Type得到构造函数，和工厂模式不返回实例不同
var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
