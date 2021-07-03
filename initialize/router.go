package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sentinelPlugin "github.com/sentinel-group/sentinel-go-adapters/gin"

	"github.com/childelins/go-gin-api/api/controller"
	"github.com/childelins/go-gin-api/global"
	"github.com/childelins/go-gin-api/middleware"
	"github.com/childelins/go-gin-api/router"
)

/*
var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.BucketRule{
		Key:          "/api/v1/auth/login",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})
*/

func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/", controller.Index)

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": "true",
		})
	})

	api := r.Group("/api/v1")
	// 配置跨域
	api.Use(middleware.Cors(), gin.Logger(), middleware.Recover())
	api.Use(middleware.Tracing())
	//api.Use(middleware.RateLimiter(methodLimiters))

	// Sentinel 会对每个 API route 进行统计，资源名称类似于 GET:/foo/:id
	// 默认的限流处理逻辑是返回 429 (Too Many Requests) 错误码，支持配置自定义的 fallback 逻辑
	api.Use(sentinelPlugin.SentinelMiddleware())
	api.Use(middleware.ContextTimeout(global.ServerConfig.ContextTimeout))
	router.NewPublicRouter(api)

	api.Use(middleware.JWT(), middleware.Admin())
	router.NewProtectedRouter(api)

	return r
}
