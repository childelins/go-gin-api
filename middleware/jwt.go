package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/childelins/go-gin-api/pkg/app"
	"github.com/childelins/go-gin-api/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			app.ToErrorResponse(c, errcode.UnauthorizedTokenError)
			c.Abort()
			return
		}

		claims, err := app.ParseToken(token)
		var ecode *errcode.Error
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				ecode = errcode.UnauthorizedTokenTimeout
			default:
				ecode = errcode.UnauthorizedTokenError
			}

			app.ToErrorResponse(c, ecode)
			c.Abort()
			return
		}

		c.Set("companyId", claims.C)
		c.Set("userId", claims.U)
		c.Next()
	}
}
