package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/childelins/go-gin-api/global"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		var parentSpan opentracing.Span
		// 直接从 c.Request.Header 中提取 span,如果没有就新建一个
		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			parentSpan = opentracing.StartSpan(c.Request.URL.Path)
		} else {
			parentSpan = opentracing.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
		}
		defer parentSpan.Finish()

		/*
			var traceID string
			var spanID string
			var spanContext = parentSpan.Context()
			switch spanContext.(type) {
			case jaeger.SpanContext:
				jaegerContext := spanContext.(jaeger.SpanContext)
				traceID = jaegerContext.TraceID().String()
				spanID = jaegerContext.SpanID().String()
			}

			c.Set("X-Trace-ID", traceID)
			c.Set("X-Span-ID", spanID)
		*/

		parentCtx := opentracing.ContextWithSpan(context.Background(), parentSpan)
		c.Set("ctx", parentCtx)

		// 设置grom context
		global.DB = global.DB.WithContext(parentCtx)
		c.Next()
	}
}
