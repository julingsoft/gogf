package redisx

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func New(groups ...string) *gredis.Redis {
	group := "default"
	if len(groups) > 0 && groups[0] != "" {
		group = groups[0]
	}
	return g.Redis(group)
}

func GetLock(ctx context.Context, key string, exp int64) (bool, error) {
	result, err := New().Set(ctx, key, 1, gredis.SetOption{
		TTLOption: gredis.TTLOption{
			EX: gconv.PtrInt64(exp), // 设置锁的有效期
		},
		NX:  true,  // 只有key不存在时才会成功设置
		Get: false, // 不需要获取原始值
	})

	if err != nil {
		return false, fmt.Errorf("failed to get lock: %v", err)
	}

	// 判断是否获取到锁
	if result.Val() == "OK" {
		return true, nil // 成功获取锁
	}

	return false, nil // 未获取到锁
}
