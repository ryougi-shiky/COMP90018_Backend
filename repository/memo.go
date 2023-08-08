package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ryougi-shiky/COMP90018_Backend/models"
	"gorm.io/gorm"
)

type MemoRepository interface {
	CreateMemo(memo *models.Memo) error
	UpdateMemo(memo *models.Memo) error
	DeleteMemo(memo *models.Memo) error
	GetUserMemo(userId *uuid.UUID) (memos []models.Memo, err error)
}

func NewMemoRepository(db *gorm.DB) MemoRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) CreateMemo(memo *models.Memo) error {
	return r.db.Create(memo).Error
}

func (r *MySQLUserRepository) UpdateMemo(memo *models.Memo) error {
	// Retrieve old version memo from database by searching memoID
	var oldMemo models.Memo
	result := r.db.Where("memo_id = ?", memo.MemoID).Find(&oldMemo).Error
	if result != nil {
		return result
	}

	// Check if both memos are from the same user
	if oldMemo.UserID != memo.UserID {
		return fmt.Errorf("Memo is from another user! Unauthorized update attempt!")
	}

	// Update memo
	updateErr := r.db.Model(&oldMemo).Updates(memo).Error
	if updateErr != nil {
		return updateErr
	}

	// Update success, no error returned
	return nil
}

func (r *MySQLUserRepository) DeleteMemo(memo *models.Memo) error {
	// Make sure the target memo is in the database
	var delMemo models.Memo
	err := r.db.Where("memo_id = ?", memo.MemoID).Find(&delMemo).Error
	if err != nil {
		return err
	}

	// Check if both memos are from the same user
	if delMemo.UserID != memo.UserID {
		return fmt.Errorf("Memo is from another user! Unauthorized update attempt!")
	}

	// Delete memo
	if err := r.db.Delete(&delMemo).Error; err != nil {
		return err
	}

	return nil
}

func (r *MySQLUserRepository) GetUserMemo(userId *uuid.UUID) (memos []models.Memo, err error) {
	// GORM find records that automatically match into the slice
	if err := r.db.Where("user_id = ?", userId).Find(&memos).Error; err != nil {
		return nil, err
	}
	return memos, nil
}
