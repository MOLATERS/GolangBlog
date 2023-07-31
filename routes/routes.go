package routes

import (
	"SimpleBlog/controller"
	"SimpleBlog/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())                              //允许跨域
	r.POST("/register", controller.Register)                        //注册
	r.POST("/login", controller.Login)                              //登录
	r.GET("/user", middleware.AuthMiddleware(), controller.GetInfo) //登录获取用户信息
	r.POST("/upload", controller.Upload)                            //用于图片传输
	r.GET("/images", controller.ShowImage)
	return r
}
