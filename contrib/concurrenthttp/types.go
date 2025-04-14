package concurrenthttp

import "time"

type Request struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    []byte
	Timeout time.Duration
}

// Result 存储每个请求的结果和错误信息
type Result struct {
	URL     string        // 请求的URL
	Body    []byte        // 响应内容
	Err     error         // 错误信息（若有的话）
	Elapsed time.Duration // 请求耗时
}
