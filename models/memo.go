package models

import (
	"github.com/google/uuid"
	"time"
)

type Memo struct {
	MemoID    uuid.UUID `gorm:"type:char(36);primary_key"`
	UserID    uuid.UUID `gorm:"type:char(36);not null"`
	Title     string    `gorm:"type:varchar(20);not null"`
	Content   string    `gorm:"type:varchar(200);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}
