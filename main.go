package main

import (
	"SimpleBlog/common"
	"SimpleBlog/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r.StaticFS("./images", http.Dir("./static/images"))
	routes.CollectRoutes(r)
	panic(r.Run(":8080"))
}
