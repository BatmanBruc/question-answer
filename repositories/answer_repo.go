package repositories

import (
	"question-answer/database"
	"question-answer/models"

	"gorm.io/gorm"
)

type AnswerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository() *AnswerRepository {
	return &AnswerRepository{
		db: database.GetDB(),
	}
}

func (r *AnswerRepository) Create(answer *models.Answer) error {
	result := r.db.Create(answer)
	return result.Error
}

func (r *AnswerRepository) GetByID(id uint) (*models.Answer, error) {
	var answer models.Answer
	result := r.db.Preload("Question").First(&answer, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &answer, nil
}

func (r *AnswerRepository) GetByQuestionID(questionID uint) ([]models.Answer, error) {
	var answers []models.Answer
	result := r.db.Where("question_id = ?", questionID).Find(&answers)
	if result.Error != nil {
		return nil, result.Error
	}
	return answers, nil
}

func (r *AnswerRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Answer{}, id)
	return result.Error
}

func (r *AnswerRepository) Update(answer *models.Answer) error {
	result := r.db.Save(answer)
	return result.Error
}

func (r *AnswerRepository) GetByUserID(userID string) ([]models.Answer, error) {
	var answers []models.Answer
	result := r.db.Where("user_id = ?", userID).Preload("Question").Find(&answers)
	if result.Error != nil {
		return nil, result.Error
	}
	return answers, nil
}
