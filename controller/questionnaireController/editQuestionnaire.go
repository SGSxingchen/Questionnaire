package questionnaireController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"questionnaire/models"
)

func EditQuestionnaire(c *gin.Context) {
	var questionnaire models.Questionnaire
	if err := c.ShouldBindJSON(&questionnaire); err != nil {
		// 如果解析请求体失败，返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
