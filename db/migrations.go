package db

import (
	"gorm.io/gorm"
	"questionnaire/models"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Questionnaire{},
		&models.Question{},
		&models.QuestionAnswer{},
		&models.UserQuestionnaire{},
	)
}
