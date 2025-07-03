package rocketmq

import (
	"context"

	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/gogf/gf/v2/frame/g"
)

type Producer struct {
	config *Config
	client rmqclient.Producer
}

func NewProducer(ctx context.Context, config *Config, opts ...rmqclient.ProducerOption) *Producer {
	var c = rmqclient.Config{
		Endpoint: config.Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    config.AccessKey,
			AccessSecret: config.AccessSecret,
		},
	}

	if config.NameSpace != "" {
		c.NameSpace = config.NameSpace
	}

	if config.ConsumerGroup != "" {
		c.ConsumerGroup = config.ConsumerGroup
	}

	opts = append(opts, rmqclient.WithTopics(config.Topic))
	producer, err := rmqclient.NewProducer(&c, opts...)
	if err != nil {
		g.Log().Error(ctx, "NewProducer error", err)
	}

	return &Producer{
		config: config,
		client: producer,
	}
}
