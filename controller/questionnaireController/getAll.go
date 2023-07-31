package questionnaireController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"questionnaire/db"
	"questionnaire/models"
)

func GetAllQuestionnaires(c *gin.Context) {
	var questionnaires []models.Questionnaire

	// 从数据库中获取全部问卷
	if err := db.DB.Find(&questionnaires).Error; err != nil {
		// 如果获取失败，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 构建响应数据
	var response struct {
		Questionnaires []struct {
			ID    uint   `json:"id"`
			Title string `json:"title"`
		} `json:"questionnaires"`
	}

	for _, questionnaire := range questionnaires {
		response.Questionnaires = append(response.Questionnaires, struct {
			ID    uint   `json:"id"`
			Title string `json:"title"`
		}{
			ID:    questionnaire.ID,
			Title: questionnaire.Title,
		})
	}

	// 返回成功响应
	c.JSON(http.StatusOK, response)
}
