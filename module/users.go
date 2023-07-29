package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:char(36);primary_key"`
	username string `gorm:"type:varchar(20);not null`
	Email string `gorm:"type:varchar(40);not null`
	Password string `gorm:"type:char(64);not null`
}