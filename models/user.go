package models

import "gorm.io/gorm"

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type User struct {
	gorm.Model
	Username string  `gorm:"type:varchar(50);uniqueIndex;not null;" validate:"required,min=3,max=50" json:"username"`
	Email    string  `gorm:"type:varchar(255);uniqueIndex;not null;" validate:"required,email,max=255" json:"email"`
	Password string  `gorm:"type:varchar(50);not null;" validate:"required,min=6,max=50" json:"password"`
	Gender   *Gender `gorm:"default:null" validate:"oneof=male female" json:"gender"`
}
