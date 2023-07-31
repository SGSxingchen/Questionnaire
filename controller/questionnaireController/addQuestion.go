package questionnaireController

import (
	"errors"
	"questionnaire/db"
	"questionnaire/models"
)

// AddQuestionToQuestionnaire 添加问题到问卷
func AddQuestionToQuestionnaire(questionnaireID uint, questionText string, questionOptions []string, questionType string) (*models.Question, error) {
	var questionnaire models.Questionnaire
	if err := db.DB.First(&questionnaire, questionnaireID).Error; err != nil {
		return nil, err
	}

	newQuestion := models.Question{
		QuestionnaireID: questionnaireID,
		Text:            questionText,
		Type:            questionType,
	}

	switch questionType {
	case "Single Choice":
		newQuestion.Options = questionOptions
	case "Multiple Choice":
		newQuestion.Options = questionOptions
	case "Subjective":
		// 没有选项
	default:
		return nil, errors.New("Invalid question type")
	}

	if err := db.DB.Create(&newQuestion).Error; err != nil {
		return nil, err
	}

	questionnaire.Questions = append(questionnaire.Questions, newQuestion)
	if err := db.DB.Save(&questionnaire).Error; err != nil {
		return nil, err
	}

	return &newQuestion, nil
}
