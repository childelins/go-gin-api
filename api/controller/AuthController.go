package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/childelins/go-gin-api/global"
	"github.com/childelins/go-gin-api/pkg/app"
	"github.com/childelins/go-gin-api/pkg/errcode"
)

type Auth struct{}

func (a *Auth) Login(c *gin.Context) {
	// 1. 根据用户名密码登录

	token, err := app.GenerateToken(app.CustomClaims{
		C: 13, // 商家id
		U: 1,  // 管理员id
	})
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		app.ToErrorResponse(c, errcode.UnauthorizedTokenGenerate)
		return
	}

	app.ToResponse(c, gin.H{
		"token": token,
	})
	return
}

func (a *Auth) Logout(c *gin.Context) {

}
