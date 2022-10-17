package day1_http_base

import (
	"context"
	"crypto/tls"
	"io"
	"mime/multipart"
	"net/url"
)

// Handler ServeHTTP处理接口
type Handler interface {
	// ServeHTTP
	// 第1个参数：可以构造针对请求的响应
	// 第2个参数：HTTP请求
	ServeHTTP(ResponseWriter, *Request)
}

type Server struct {
	Addr    string
	Handler Handler
}

// ListenAndServe
// addr：监听地址
// handler：所有实现ServeHTTP接口的实例，所有的HTTP请求都交给该实例处理
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

type Header map[string][]string

type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

// Request HTTP请求
type Request struct {
	URL    *url.URL
	Header Header
	Body   io.ReadCloser

	Method     string
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0

	GetBody          func() (io.ReadCloser, error)
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Host             string
	Form             url.Values
	PostForm         url.Values

	MultipartForm *multipart.Form
	Trailer       Header
	RemoteAddr    string
	RequestURI    string
	TLS           *tls.ConnectionState
	Response      *Response
	ctx           context.Context
}

type Response struct {
	Status           string // e.g. "200 OK"
	StatusCode       int    // e.g. 200
	Proto            string // e.g. "HTTP/1.0"
	ProtoMajor       int    // e.g. 1
	ProtoMinor       int    // e.g. 0
	Header           Header
	Body             io.ReadCloser
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Uncompressed     bool
	Trailer          Header
	Request          *Request
	TLS              *tls.ConnectionState
}
