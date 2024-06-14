package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CheckCredentialUser(user User) (result User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(database *gorm.DB) UserRepository {
	return &userRepository{
		database,
	}
}

func (r *userRepository) CheckCredentialUser(user User) (result User, err error) {
	if err = r.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&result).Error; err != nil {
		return
	}

	return
}
