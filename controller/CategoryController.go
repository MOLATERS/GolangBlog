package controller

import (
	"SimpleBlog/common"
	"SimpleBlog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* controller/CategoryController.go */
// SearchCategory 查询分类

func SearchCategory(c *gin.Context) {
	db := common.GetDB()
	var categories []model.Category
	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":       200,
			"categories": nil,
			"msg":        "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":       200,
		"categories": categories,
		"msg":        "查找成功",
	})
}

// SearchCategoryName 查询分类名
func SearchCategoryName(c *gin.Context) {
	db := common.GetDB()
	var category model.Category
	// 获取path中的分类id
	categoryId := c.Params.ByName("id")
	sql := "SELECT * FROM categories WHERE id = ?"
	//db.Where("id = ?", categoryId).First(&category).Error
	if err := db.Raw(sql, categoryId).First(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"data": "",
			"msg":  "查找失败",
		})
		return
	}
	//if err := db.Take(&category,model.Category{ID: (unint)(categoryId)},nil).Error; err{
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"code": 500,
	//			"data": "",
	//			"msg":  "查找失败",
	//		})
	//		return
	//}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"name": category.CategoryName,
		},
		"msg": "查找成功",
	})
}
