package models

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey"`
	QuestionID uint   `json:"question_id" gorm:"not null;index"`
	UserID     uint   `json:"user_id" gorm:"not null;index"`
	Text       string `json:"text" gorm:"type:text;not null"`

	Question Question `json:"question,omitempty" gorm:"foreignKey:QuestionID"`
	User     User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type CreateAnswerRequest struct {
	Text string `json:"text" validate:"required,min=1,max=5000"`
}
