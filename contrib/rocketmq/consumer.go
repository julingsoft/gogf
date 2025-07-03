package rocketmq

import (
	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

func NewConsumer(config *Config, opts ...rmqclient.SimpleConsumerOption) (rmqclient.SimpleConsumer, error) {
	return rmqclient.NewSimpleConsumer(&rmqclient.Config{
		Endpoint:      config.Endpoint,
		NameSpace:     config.NameSpace,
		ConsumerGroup: config.ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    config.AccessKey,
			AccessSecret: config.AccessSecret,
		},
	},
		opts...,
	)
}
