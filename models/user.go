package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primary_key"`
	Username string    `gorm:"type:varchar(20);not null;unique"`
	Email    string    `gorm:"type:varchar(40);not null;unique"`
	Password string    `gorm:"type:char(64);not null"`
}
