package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID      uint     `json:"id" gorm:"primaryKey"`
	Text    string   `json:"text" gorm:"type:text;not null"`
	Answers []Answer `json:"answers,omitempty" gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE;"`
}

type CreateQuestionRequest struct {
	Text string `json:"text"`
}
