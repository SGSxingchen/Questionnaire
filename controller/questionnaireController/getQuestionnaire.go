package questionnaireController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"questionnaire/db"
	"questionnaire/models"
)

func GetQuestionnaireByID(c *gin.Context) {
	// 从URL参数中获取问卷ID
	questionnaireID := c.Param("id")

	var questionnaire models.Questionnaire

	// 从数据库中获取指定ID的问卷
	if err := db.DB.First(&questionnaire, questionnaireID).Error; err != nil {
		// 如果获取失败，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 查询关联的问题
	var questions []models.Question
	if err := db.DB.Where("questionnaire_id = ?", questionnaireID).Find(&questions).Error; err != nil {
		// 如果查询问题失败，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将问题添加到问卷对象中
	questionnaire.Questions = questions

	// 返回成功响应
	c.JSON(http.StatusOK, questionnaire)
}
