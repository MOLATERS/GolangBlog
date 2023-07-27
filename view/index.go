package views

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func IsODD(num int) bool {
	return num%2 == 0
}
func date(layout string) string {
	return time.Now().Format(layout)
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "第一个尝试的博客"
	indexData.Desc = "现在是入门教程"
	t := template.New("index.html")
	Path := config.Cfg.System.CurrentDir
	home := Path + "/template/home.html"
	header := Path + "/template/layout/header.html"
	footer := Path + "/template/layout/footer.html"
	personal := Path + "/template/layout/personal.html"
	post := Path + "/template/layout/post-list.html"
	pagination := Path + "/template/layout/pagination.html"
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": date})
	t, err := t.ParseFiles(Path+"/template/index.html", home, footer, header, personal, post, pagination)
	if err != nil {
		log.Println("解析模板出错", err)
	}
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	t.Execute(w, hr)
}
