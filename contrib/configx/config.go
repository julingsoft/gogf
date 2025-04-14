package configx

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
)

type Config struct {
	App     App
	Logging Logging
	OSS     OSS
	Redis   Redis
}

type App struct {
	Name  string `json:"name"`
	Debug bool   `json:"debug"`
}

type Logging struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	ProjectName     string
	LogStoreName    string
}

type OSS struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	RegionName      string
	BucketName      string
}
type Redis struct {
	Default RedisConfig
	Cache   RedisConfig
}

type RedisConfig struct {
	Address     string
	Pass        string
	Db          int
	IdleTimeout int
}

var configData *Config

func New(ctx context.Context) *Config {
	if configData != nil {
		return configData
	}

	var app App
	appConf := g.Config().MustGet(ctx, "app")
	err := json.Unmarshal(appConf.Bytes(), &app)
	if err != nil {
		g.Log().Error(ctx, "app config error", err)
		return nil
	}

	var logging Logging
	loggingConf := g.Config().MustGet(ctx, "service.logging")
	err = json.Unmarshal(loggingConf.Bytes(), &logging)
	if err != nil {
		g.Log().Error(ctx, "sls config error", err)
		return nil
	}

	var oss OSS
	ossConf := g.Config().MustGet(ctx, "service.oss")
	err = json.Unmarshal(ossConf.Bytes(), &oss)
	if err != nil {
		g.Log().Error(ctx, "oss config error", err)
		return nil
	}

	var redis Redis
	redisConf := g.Config().MustGet(ctx, "redis")
	err = json.Unmarshal(redisConf.Bytes(), &redis)
	if err != nil {
		g.Log().Error(ctx, "redis config error", err)
		return nil
	}

	configData = &Config{
		App:     app,
		Logging: logging,
		OSS:     oss,
		Redis:   redis,
	}

	return configData
}
