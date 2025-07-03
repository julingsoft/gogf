package rocketmq

import "time"

type Config struct {
	Topic         string        `json:"topic"`
	Endpoint      string        `json:"endpoint"`
	NameSpace     string        `json:"nameSpace"`
	ConsumerGroup string        `json:"consumerGroup"`
	AccessKey     string        `json:"accessKey"`
	AccessSecret  string        `json:"accessSecret"`
	SecurityToken string        `json:"securityToken"`
	DelayTime     time.Duration `json:"delayTime"`
}
