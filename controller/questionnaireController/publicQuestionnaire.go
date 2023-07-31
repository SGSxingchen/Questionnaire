package questionnaireController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"questionnaire/db"
	"questionnaire/models"
)

func IsPublic(c *gin.Context) {
	// 从URL参数中获取问卷ID
	questionnaireID := c.Param("id")

	var questionnaire models.Questionnaire

	// 从数据库中获取指定ID的问卷
	if err := db.DB.First(&questionnaire, questionnaireID).Error; err != nil {
		// 如果获取失败，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, questionnaire.IsPublic)
}

func setPublic(c *gin.Context) {
	questionnaireID := c.Param("id")
	publicStatus := c.Param("is_public")
	var questionnaire models.Questionnaire
	// 从数据库中获取指定ID的问卷
	if err := db.DB.First(&questionnaire, questionnaireID).Error; err != nil {
		// 如果获取失败，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if publicStatus == "true" {
		questionnaire.IsPublic = true
	} else {
		questionnaire.IsPublic = false
	}
	c.JSON(http.StatusOK, "OK")
}
