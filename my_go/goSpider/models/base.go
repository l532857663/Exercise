package models

import "net/http"

// 数据请求的类型
type Request struct {
	// HTTP请求
	httpReq *http.Request
	// 请求的深度
	depth uint32
}

// 创建一个新的请求实例
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq, depth: depth}
}

// 用于获取的请求的HTTP信息
func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

// 用于获取请求的深度
func (req *Request) Depth() uint32 {
	return req.depth
}
