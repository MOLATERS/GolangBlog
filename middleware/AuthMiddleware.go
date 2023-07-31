package middleware

import (
	"SimpleBlog/common"
	"SimpleBlog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		// 获取Header
		tokenString :=c.Request.Header.Get("Authorization")
		// token 为空
		if tokenString == ""{
			c.JSON(http.StatusOK,gin.H{
				"code" : 401,
				"msg" : "权限不足",
			})
		c.Abort()
			return
		}
		//非法token
		if tokenString == "" || len(tokenString) < 7 ||!strings.HasPrefix(tokenString,"Bearer"){
			c.JSON(http.StatusOK,gin.H{
				"code" : 401,
				"msg" : "权限不足",
			})
			c.Abort()
			return
		}
		//提取Token的有效成分
		tokenString = tokenString[7:]
		token,claims,err :=common.ParseToken(tokenString)
		if err!=nil || !token.Valid{
			c.JSON(http.StatusOK,gin.H{
				"code" : 401,
				"msg" : "权限不足",
			})
			c.Abort()
			return
		}
		//获取claims中的userId
		userId := claims.UserId
		DB := common.DB
		var user model.User
		DB.Where("id =?",userId).First(&user)
		c.Set("user",user)
		c.Next()
	}
}



