package ossx

import (
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/gogf/gf/v2/frame/g"
	"io"
)

type Config struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	RegionName      string `json:"regionName"`
	BucketName      string `json:"bucketName"`
}

type OSS struct {
	config *Config
	client *oss.Client
}

func New(config *Config) *OSS {
	provider := credentials.NewStaticCredentialsProvider(config.AccessKeyId, config.AccessKeySecret)
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(provider).
		WithRegion(config.RegionName)

	return &OSS{
		config: config,
		client: oss.NewClient(cfg),
	}
}

func (o *OSS) Client() *oss.Client {
	return o.client
}

func (o *OSS) PutObject(ctx context.Context, objectName string, body io.Reader) (*oss.PutObjectResult, error) {
	putRequest := &oss.PutObjectRequest{
		Bucket:       oss.Ptr(o.config.BucketName), // 存储空间名称
		Key:          oss.Ptr(objectName),          // 对象名称
		Body:         body,                         // 对象内容
		StorageClass: oss.StorageClassStandard,     // 指定对象的存储类型为标准存储
		Acl:          oss.ObjectACLPrivate,         // 指定对象的访问权限为私有访问
	}

	return o.Client().PutObject(ctx, putRequest)
}

func (o *OSS) PutObjectFromFile(ctx context.Context, objectName string, filePath string) (*oss.PutObjectResult, error) {
	putRequest := &oss.PutObjectRequest{
		Bucket:       oss.Ptr(o.config.BucketName), // 存储空间名称
		Key:          oss.Ptr(objectName),          // 对象名称
		StorageClass: oss.StorageClassStandard,     // 指定对象的存储类型为标准存储
		Acl:          oss.ObjectACLPrivate,         // 指定对象的访问权限为私有访问
	}

	return o.Client().PutObjectFromFile(ctx, putRequest, filePath)
}

func (o *OSS) GetObject(ctx context.Context, objectName string) (*oss.GetObjectResult, error) {
	getRequest := &oss.GetObjectRequest{
		Bucket: oss.Ptr(o.config.BucketName), // 存储空间名称
		Key:    oss.Ptr(objectName),          // 对象名称
	}

	return o.Client().GetObject(ctx, getRequest)
}

func (o *OSS) MustGetObject(ctx context.Context, objectName string) string {
	result, err := o.GetObject(ctx, objectName)
	if err != nil {
		g.Log().Error(ctx, err, "[ossx] GetObject", objectName)
		return ""
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
	if err != nil {
		g.Log().Error(ctx, err, "[ossx] ReadAll", objectName)
		return ""
	}

	return string(data)
}
