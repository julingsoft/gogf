package rocketmq

import (
	"context"

	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gogf/gf/v2/frame/g"
)

func (p *Producer) Async(ctx context.Context, tag string, keys []string, messages []*rmqclient.Message) error {
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

		// send message in async
		p.client.SendAsync(ctx, message, func(ctx context.Context, resp []*rmqclient.SendReceipt, err error) {
			if err != nil {
				g.Log().Error(ctx, err)
			}
			for i := 0; i < len(resp); i++ {
				g.Log().Infof(ctx, "send message success. %#v\n", resp[i])
			}
		})
	}

	return nil
}
