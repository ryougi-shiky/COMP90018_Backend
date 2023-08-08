package repository

import (
	"github.com/ryougi-shiky/COMP90018_Backend/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

// Create a new repository object. Connect to the database.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) RegisterUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *MySQLUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
