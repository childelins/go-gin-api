package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type ILimiter interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...BucketRule) ILimiter
}

type Limiter struct {
	buckets map[string]*ratelimit.Bucket
}

type BucketRule struct {
	Key          string        // 自定义键值对名称
	FillInterval time.Duration // 间隔多长时间放N个令牌
	Capacity     int64         // 令牌桶的容量
	Quantum      int64         // 每次到达间隔时间后所放的具体令牌数量
}
