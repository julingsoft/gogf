package rocketmq

import (
	"context"
	"time"

	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gogf/gf/v2/frame/g"
)

func (p *Producer) Delay(ctx context.Context, tag string, keys []string, messages []*rmqclient.Message) error {
	// start producer
	err := p.client.Start()
	if err != nil {
		return err
	}

	// graceful stop producer
	defer func(producer rmqclient.Producer) {
		err := producer.GracefulStop()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(p.client)

	for _, message := range messages {
		message.SetTag(tag)
		message.SetKeys(keys...)
		message.SetDelayTimestamp(time.Now().Add(p.config.DelayTime))

		// send message in async
		resp, err := p.client.Send(ctx, message)
		if err != nil {
			g.Log().Error(ctx, err)
		}

		for i := 0; i < len(resp); i++ {
			g.Log().Infof(ctx, "send message success. %#v\n", resp[i])
		}
	}

	return nil
}
