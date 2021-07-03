package app

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/childelins/go-gin-api/global"
)

func Validate(c *gin.Context, v interface{}) (bool, error) {
	if err := c.ShouldBind(v); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			var errMsg []string
			for _, v := range errs.Translate(global.Trans) {
				errMsg = append(errMsg, v)
			}
			return false, fmt.Errorf("%v", strings.Join(errMsg, ","))
		}

		return false, err
	}

	return true, nil
}
