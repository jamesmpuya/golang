package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int       `json:"id" gorm:"primary_key"`
	Uid      uuid.UUID `json:"uid" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name" gorm:"primary_key"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
