package router

import (
	"awesomeProject/api"
	"awesomeProject/view"
	"net/http"
)

// 访问模板的时候需要将嵌套的模板都进行解析

func Router() {
	// 1.页面 viewer 2.数据 json api 3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
