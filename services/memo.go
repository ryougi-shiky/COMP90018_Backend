package services

import (
	"github.com/google/uuid"
	"github.com/ryougi-shiky/COMP90018_Backend/models"
	"github.com/ryougi-shiky/COMP90018_Backend/repository"
)

type MemoService interface {
	CreateMemo(memo *models.Memo) error
	UpdateMemo(memo *models.Memo) error
	DeleteMemo(memo *models.Memo) error
	GetUserMemo(userId *uuid.UUID) (memos []models.Memo, err error)
}

type MemoServiceImpl struct {
	MemoRepository repository.MemoRepository
}

func (s *MemoServiceImpl) CreateMemo(memo *models.Memo) error {
	return s.MemoRepository.CreateMemo(memo)
}

func (s *MemoServiceImpl) UpdateMemo(memo *models.Memo) error {
	return s.MemoRepository.UpdateMemo(memo)
}

func (s *MemoServiceImpl) DeleteMemo(memo *models.Memo) error {
	return s.MemoRepository.DeleteMemo(memo)
}

func (s *MemoServiceImpl) GetUserMemo(userId *uuid.UUID) (memos []models.Memo, err error) {
	return s.MemoRepository.GetUserMemo(userId)
}

func NewMemoService(memoRepo repository.MemoRepository) MemoService {
	return &MemoServiceImpl{
		MemoRepository: memoRepo,
	}
}
