package ossx

import (
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/julingsoft/gogf/contrib/configx"
)

func New(ctx context.Context) *oss.Client {
	c := configx.New(ctx)

	provider := credentials.NewStaticCredentialsProvider(c.OSS.AccessKeyID, c.OSS.AccessKeySecret)
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(provider).
		WithRegion(c.OSS.RegionName)

	return oss.NewClient(cfg)
}
