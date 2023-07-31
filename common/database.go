package common

import(
	"SimpleBlog/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// initDB 数据库初始化

func InitDB() *gorm.DB {
	driverName := "mysql"
	user :="root"
	password := "123456"
	host := "localhost"
	port := "3306"
	database := "blog"
	charset :="utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",user,password,host,port,database,charset)
	db,err:=gorm.Open(driverName,args)
	if err != nil {
		panic("fail to open database: "+ err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db

}
func GetDB() *gorm.DB{
	return DB
}
