package questionnaireController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"questionnaire/db"
	"questionnaire/models"
)

func DeleteQuestionnaire(c *gin.Context) {
	questionnaireID := c.Param("id")

	var questionnaire models.Questionnaire
	if err := db.DB.First(&questionnaire, questionnaireID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Questionnaire not found"})
		return
	}

	if err := db.DB.Delete(&questionnaire).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Questionnaire deleted successfully"})
}
