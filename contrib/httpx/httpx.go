package httpx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/julingsoft/gogf/contrib/logx"
	"net/http"
	"time"
)

type HttpClient struct {
	BaseUrl string
	timeout time.Duration // 请求超时 单位毫秒
}

func New(baseUrl string) *HttpClient {
	return &HttpClient{
		BaseUrl: baseUrl,
		timeout: 5 * time.Second,
	}
}

func (c *HttpClient) SetTimeout(timeout time.Duration) *HttpClient {
	c.timeout = timeout
	return c
}

func (c *HttpClient) Get(ctx context.Context, url string) ([]byte, error) {
	startTime := time.Now()

	url = gstr.TrimRight(c.BaseUrl, "/") + "/" + gstr.TrimLeft(url, "/")
	r, err := g.Client().SetTimeout(c.timeout).Get(ctx, url)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	duration := time.Since(startTime).Milliseconds()
	g.Log().Info(ctx, logx.LogData{
		Method:   http.MethodGet,
		Url:      url,
		Response: r.ReadAllString(),
		Status:   r.Response.StatusCode,
		Duration: duration,
		LogTime:  time.Now(),
	})

	return r.ReadAll(), nil
}

func (c *HttpClient) Post(ctx context.Context, url string, data string) ([]byte, error) {
	startTime := time.Now()

	url = gstr.TrimRight(c.BaseUrl, "/") + "/" + gstr.TrimLeft(url, "/")
	r, err := g.Client().SetTimeout(c.timeout).Post(ctx, url, data)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	duration := time.Since(startTime).Milliseconds()
	g.Log().Info(ctx, logx.LogData{
		Method:   http.MethodPost,
		Url:      url,
		Request:  data,
		Response: r.ReadAllString(),
		Status:   r.Response.StatusCode,
		Duration: duration,
		LogTime:  time.Now(),
	})

	return r.ReadAll(), nil
}
