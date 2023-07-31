package questionnaireController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"questionnaire/db"
	"questionnaire/models"
)

type NewQuestion2 struct {
	ID      uint     `json:"id" binding:"required"`
	Text    string   `json:"text" binding:"required"`
	Options []string `json:"options" binding:"required"`
	Type    string   `json:"type" binding:"required"`
}

func AddMultipleQuestion(c *gin.Context) {
	// 解析请求体中的问题数据
	var newQuestion NewQuestion2
	if err := c.ShouldBindJSON(&newQuestion); err != nil {
		// 如果解析请求体失败，返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newQuestion.Type != "Multiple" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type非法"})
		return
	}

	// 创建问题对象
	question := models.Question{
		QuestionnaireID: newQuestion.ID,
		Text:            newQuestion.Text,
		Options:         newQuestion.Options,
		Type:            newQuestion.Type,
	}

	// 将问题保存到数据库
	if err := db.DB.Create(&question).Error; err != nil {
		// 如果保存失败，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusCreated, question)
}
