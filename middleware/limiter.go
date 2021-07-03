package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/childelins/go-gin-api/pkg/app"
	"github.com/childelins/go-gin-api/pkg/errcode"
	"github.com/childelins/go-gin-api/pkg/limiter"
)

func RateLimiter(l limiter.ILimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				app.ToErrorResponse(c, errcode.TooManyRequests)
				c.Abort()
			}
		}

		c.Next()
	}

}
