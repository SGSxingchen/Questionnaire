package models

import (
	"gorm.io/gorm"
)

type Questionnaire struct {
	gorm.Model
	Title     string     `json:"title"`
	IsPublic  bool       `json:"is_public"`
	Questions []Question `gorm:"constraint:OnDelete:CASCADE" json:"questions"`
}

type Question struct {
	gorm.Model
	QuestionnaireID uint     `json:"-"`
	Text            string   `json:"text"`
	Options         []string `gorm:"type:varchar(255);" json:"options"`
	Type            string   `json:"type"`
}
type QuestionAnswer struct {
	gorm.Model
	UserID          uint     `json:"-"`
	QuestionnaireID uint     `json:"-"`
	Text            string   `json:"text"`
	Answer          []string `gorm:"type:varchar(255);" json:"options"`
	Type            string   `json:"type"`
}
type UserQuestionnaire struct {
	gorm.Model
	UserID          uint             `json:"user_id"`
	QuestionnaireID uint             `json:"questionnaire_id"`
	QuestionAnswer  []QuestionAnswer `gorm:"constraint:OnDelete:CASCADE" json:"question_answer"`
}
