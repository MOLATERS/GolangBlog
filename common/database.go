package common

import (
	"SimpleBlog/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// initDB 数据库初始化

func InitDB() *gorm.DB {
	driverName := "mysql"
	user := "root"
	password := "123456"
	host := "localhost"
	port := "13306"
	database := "blog"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", user, password, host, port, database, charset)
	//args+=fmt.Sprintf(`UNIX_TIMESTAMP("%s") - UNIX_TIMESTAMP(created_at)`, time.Now().Format("2006-01-02 15:04:05"))
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("fail to open database: " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}
