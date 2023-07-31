package controller

import (
	"SimpleBlog/common"
	"SimpleBlog/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context){
	db := common.GetDB()
	//获取参数

	var requestUser model.User
	c.Bind(&requestUser)
	userName := requestUser.UserName
	phoneNumber := requestUser.PhoneNumber
	passWord := requestUser.Password

	//数据验证

	var user model.User
	db.Where("phone_number = ?",phoneNumber).First(&user)
	if user.ID !=0 {
		c.JSON(http.StatusOK,gin.H{
			"code" : 422,
			"msg" : "用户已存在",
		})
		return
	}

	//密码加密

	hashedPassword,_ := bcrypt.GenerateFromPassword([]byte(passWord),bcrypt.DefaultCost)

	//创建用户

	newUser := model.User{
		UserName :   userName,
		PhoneNumber: phoneNumber,
		Password: string(hashedPassword),
		Avatar:      "/images/default_avatar.png",
		Collects:    model.Array{},
		Following:   model.Array{},
		Fans:        0,
	}
	db.Create(&newUser)

	//返回结果

	c.JSON(http.StatusOK,gin.H{
		"code" : 200,
		"msg" : "注册成功",
	})
}

func Login(c *gin.Context){
	db := common.GetDB()

	//获取参数

	var requestUser model.User
	c.Bind(&requestUser)
	phoneNumber := requestUser.PhoneNumber
	passWord := requestUser.Password

	//数据验证

	var user model.User
	db.Where("phone_number = ?",phoneNumber).First(&user)
	if user.ID==0{
		c.JSON(http.StatusOK,gin.H{
			"code" : 422,
			"msg" : "用户不存在",
		})
	 return
	}

	//判断密码是否正确

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(passWord));err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code" : 422,
			"msg" : "密码错误",
		})
	return
	}

	//发放token

	token,err := common.RleaseToken(user)
	if err !=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"data":gin.H{
				"token" :token,
			},
			"msg" : "登陆成功",
		})
		return
	}

}

//验证是否有效

func GetInfo(c *gin.Context) {
	// 获取上下文中的用户信息
	user, _ := c.Get("user")
	// 返回用户信息
	//response.Success(c, gin.H{"id": user.(model.User).ID, "avatar": user.(model.User).Avatar}, "登录获取信息成功")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"id": user.(model.User).ID, "avatar": user.(model.User).Avatar},
		"msg":  "登录获取信息成功",
	})
}
