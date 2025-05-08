package cachex

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

func New(groups ...string) (cache *gcache.Cache) {
	group := "cache"
	if len(groups) > 0 && groups[0] == "" {
		group = groups[0]
	}
	return gcache.NewWithAdapter(gcache.NewAdapterRedis(g.Redis(group)))
}
