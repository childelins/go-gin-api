package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() ILimiter {
	return MethodLimiter{&Limiter{buckets: make(map[string]*ratelimit.Bucket)}}
}

func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}

	return uri[:index]
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.buckets[key]
	return bucket, ok
}

func (l MethodLimiter) AddBuckets(rules ...BucketRule) ILimiter {
	for _, rule := range rules {
		if _, ok := l.buckets[rule.Key]; !ok {
			l.buckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}

	return l
}
