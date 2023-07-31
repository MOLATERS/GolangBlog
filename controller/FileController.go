package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func Upload(c *gin.Context){

	file,header,err :=  c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"code" : 500,
			"msg" : "格式错误",
		})
		return
	}
	filename := header.Filename
	exit := path.Ext(filename)
	//用上传时间作为文件名
	name := "image_" + time.Now().Format("20060102150405")
	newFilename := name + exit
	Filepath := "./images"
	err = os.MkdirAll(Filepath, os.ModePerm)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"code" : 500,
			"msg" : "创建错误 " + err.Error(),
		})
	}
	out, err := os.Create(Filepath + newFilename)
	//out,err := os.OpenFile(Filepath,os.O_RDWR|os.O_CREATE,0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":500,
			"msg":"创建错误 : " + err.Error(),
		})
		return
	}
	defer file.Close()
	_,err = io.Copy(out,file)
	//fmt.Println("%+v",file)
	//fmt.Println("%+v",out)
	defer out.Close()
	//fmt.Println("%+v",file)
	//fmt.Println("%+v",out)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":500,
			"msg":"复制错误",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":gin.H{
			"filePath":"images/"+newFilename,
		},
		"msg":"上传成功",
	})
}

func ShowImage(c *gin.Context){
	imageName := c.Query("imageName")
	c.File(imageName)
}