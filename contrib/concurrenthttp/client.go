package concurrenthttp

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Client 配置并发HTTP客户端
type Client struct {
	Concurrency int
	Timeout     time.Duration
}

// NewClient 创建新的并发HTTP客户端
func NewClient(concurrency int, timeout time.Duration) *Client {
	return &Client{
		Concurrency: concurrency,
		Timeout:     timeout,
	}
}

// Fetch 发送并发请求并返回结果
func (c *Client) Fetch(urls []Request) ([]Result, error) {
	var (
		results = make(chan Result, len(urls))
		wg      sync.WaitGroup
	)

	// 初始化工作池
	taskChan := make(chan Request, c.Concurrency)
	wg.Add(c.Concurrency)
	for i := 0; i < c.Concurrency; i++ {
		go c.worker(taskChan, results, &wg)
	}

	// 发送任务到通道
	for _, url := range urls {
		taskChan <- url
	}
	close(taskChan)

	// 等待所有任务完成并关闭结果通道
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集结果到切片
	finalResults := make([]Result, 0, len(urls))
	for res := range results {
		finalResults = append(finalResults, res)
	}

	return finalResults, nil
}

// worker 处理单个请求的工作函数
func (c *Client) worker(taskChan <-chan Request, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Timeout: c.Timeout,
	}

	for request := range taskChan {
		method := "GET"
		if request.Method != "" {
			method = request.Method
		}

		var reqBody io.Reader
		if request.Body != nil {
			reqBody = strings.NewReader(string(request.Body))
		}

		start := time.Now()
		req, _ := http.NewRequest(method, request.URL, reqBody)

		if request.Headers != nil {
			for key, value := range request.Headers {
				req.Header.Set(key, value)
			}
		}

		if request.Timeout != 0 {
			client.Timeout = request.Timeout
		}

		resp, err := client.Do(req)
		elapsed := time.Since(start)

		var body []byte
		if err == nil {
			defer resp.Body.Close()
			body, _ = io.ReadAll(resp.Body)
		}

		results <- Result{
			URL:     request.URL,
			Body:    body,
			Err:     err,
			Elapsed: elapsed,
		}
	}
}
