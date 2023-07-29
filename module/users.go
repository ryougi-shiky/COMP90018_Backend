package models

import {
	"github.com/google/uuid"
}

type User struct {
	ID uuid.UUID `gorm:"type:char(36);primary_key"`
}