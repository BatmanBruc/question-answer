package repositories

import (
	"question-answer/database"
	"question-answer/models"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{
		db: database.GetDB(),
	}
}

func (r *QuestionRepository) Create(question *models.Question) error {
	result := r.db.Create(question)
	return result.Error
}

func (r *QuestionRepository) GetByID(id uint) (*models.Question, error) {
	var question models.Question
	result := r.db.First(&question, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &question, nil
}

func (r *QuestionRepository) GetWithAnswers(id uint) (*models.Question, error) {
	var question models.Question
	result := r.db.Preload("Answers").First(&question, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &question, nil
}

func (r *QuestionRepository) GetAll() ([]models.Question, error) {
	var questions []models.Question
	result := r.db.Find(&questions)
	if result.Error != nil {
		return nil, result.Error
	}
	return questions, nil
}

func (r *QuestionRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("question_id = ?", id).
			Delete(&models.Answer{}).Error; err != nil {
			return err
		}
		return tx.Unscoped().Delete(&models.Question{}, id).Error
	})
}

func (r *QuestionRepository) Update(question *models.Question) error {
	result := r.db.Save(question)
	return result.Error
}
