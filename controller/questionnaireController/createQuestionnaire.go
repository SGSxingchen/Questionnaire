package questionnaireController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"questionnaire/db"
	"questionnaire/models"
)

//	func CreateQuestionnaire(title string) (*models.Questionnaire, error) {
//		questionnaire := models.Questionnaire{
//			Title: title,
//		}
//
//		if err := db.DB.Create(&questionnaire).Error; err != nil {
//			return nil, err
//		}
//
//		return &questionnaire, nil
//	}
func CreateQuestionnaire(c *gin.Context) {
	// 解析请求参数
	var req struct {
		Title string `json:"title" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		// 如果请求参数解析失败，返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建一个空白问卷对象
	questionnaire := models.Questionnaire{
		Title:    req.Title,
		IsPublic: false,
	}

	// 将问卷对象保存到数据库中
	if err := db.DB.Create(&questionnaire).Error; err != nil {
		// 如果保存失败，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"questionnaire": questionnaire})
}
