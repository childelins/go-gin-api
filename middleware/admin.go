package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/childelins/go-gin-api/pkg/app"
	"github.com/childelins/go-gin-api/pkg/errcode"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Get("userId")
		if !ok || userId != 1 {
			app.ToErrorResponse(c, errcode.AuthorizationForbidden)
			c.Abort()
			return
		}
		c.Next()
	}
}
