package router

import (
	"github.com/gin-gonic/gin"
	"go_code/blog/admin/controllers"
	"go_code/blog/admin/middleware"
	"net/http"
)

func InitRouter(r *gin.Engine) *gin.Engine{
	//定义路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	loginGroup := r.Group("login")
	{
		login := &controllers.Login{}
		loginGroup.GET("/captcha",login.Captcha)
		loginGroup.GET("/showLogin",login.ShowLogin)
		loginGroup.POST("/toLogin",login.ToLogin)
	}
	//    .Use(middleware.IsLogin())
	adminGroup := r.Group("/admin").Use(middleware.IsLogin())
	{
		admin := &controllers.Index{}
		adminGroup.GET("/index",admin.ShowIndex)
		adminGroup.POST("/outLogin",admin.OutLogin)
		adminGroup.GET("/cateIndex",admin.CateIndex)
		adminGroup.GET("/showCateAdd",admin.ShowAddCate)
	}
	return r
}