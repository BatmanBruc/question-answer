package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      uint     `json:"id" gorm:"primaryKey"`
	Token   string   `json:"token" gorm:"uniqueIndex;not null"`
	Answers []Answer `json:"answers,omitempty" gorm:"foreignKey:UserID"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
