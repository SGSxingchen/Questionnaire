package questionnaireController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"questionnaire/db"
	"questionnaire/models"
)

func SubmitQuestionnaire(c *gin.Context) {
	var userQuestionnaire models.UserQuestionnaire

	if err := c.ShouldBindJSON(&userQuestionnaire); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查问卷是否存在
	var questionnaire models.Questionnaire
	if err := db.DB.First(&questionnaire, userQuestionnaire.QuestionnaireID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Questionnaire not found"})
		return
	}

	// 保存到数据库中
	if err := db.DB.Create(&userQuestionnaire).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Questionnaire response submitted successfully"})
}
