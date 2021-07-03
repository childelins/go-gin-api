package router

import (
	"github.com/gin-gonic/gin"

	"github.com/childelins/go-gin-api/api/controller"
)

func NewPublicRouter(r *gin.RouterGroup) {

	auth := &controller.Auth{}
	authApi := r.Group("auth")
	{
		authApi.POST("login", auth.Login)
		authApi.POST("logout", auth.Logout)
	}
}

// 接口请求重定向问题 https://github.com/gin-gonic/gin/issues/1004
func NewProtectedRouter(r *gin.RouterGroup) {
	lecturer := &controller.Lecturer{}
	lecturerApi := r.Group("lecturers")
	{
		lecturerApi.GET("", lecturer.List)
		lecturerApi.POST("", lecturer.Create)
		lecturerApi.PUT(":id", lecturer.Update)
		lecturerApi.DELETE(":id", lecturer.Delete)
	}
}
