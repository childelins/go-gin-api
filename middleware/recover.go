package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/childelins/go-gin-api/global"
	"github.com/childelins/go-gin-api/pkg/app"
	"github.com/childelins/go-gin-api/pkg/errcode"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.Error(err)

				// TODO: 可加入其他报警处理

				app.ToErrorResponse(c, errcode.ServerError)
				c.Abort()
			}
		}()

		c.Next()
	}
}
