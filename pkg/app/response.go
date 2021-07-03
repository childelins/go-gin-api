package app

import (
	"net/http"

	"github.com/childelins/go-gin-api/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func ToErrorResponse(c *gin.Context, err *errcode.Error) {
	c.JSON(err.StatusCode(), gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	})
}

func ToResponse(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}
