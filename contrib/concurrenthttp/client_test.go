package concurrenthttp

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	// 创建客户端配置
	client := NewClient(3, 100*time.Millisecond)

	// 待请求的URL列表
	urls := []Request{
		{
			URL:    "https://www.baidu.com",
			Method: "POST",
		},
		{
			URL: "https://www.qq.com",
		},
		{
			URL: "https://www.x.com",
		},
	}

	// 执行并发请求
	results, err := client.Fetch(urls)
	if err != nil {
		panic(err)
	}

	// 处理结果
	for _, res := range results {
		if res.Err != nil {
			fmt.Printf("Error fetching %s: %v\n", res.URL, res.Err)
			continue
		}
		fmt.Printf("Result from %s (%v):\n%v\n", res.URL, res.Elapsed, string(res.Body))
	}
}
