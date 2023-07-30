package services

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/ryougi-shiky/COMP90018_Backend/models"
	"github.com/ryougi-shiky/COMP90018_Backend/repository"
)

type UserService interface {
	RegisterUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (s *UserServiceImpl) RegisterUser(user *models.User) error  {
	// sha256 hashing the password
	hasher := sha256.New()
	hasher.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hasher.Sum(nil))

	return s.UserRepository.CreateUser(user)
}

func (s *UserServiceImpl) GetUserByEmail(email string) (*models.User, error) {
	return s.UserRepository.GetUserByEmail(email)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepo,
	}
}