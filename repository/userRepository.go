package repository

import (
	"github.com/Darari17/golang-e-commerce/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	FindUserById(userId uint) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

// CreateUser implements IUserRepository.
func (u *userRepository) CreateUser(user *model.User) (*model.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindByEmail implements IUserRepository.
func (u *userRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindById implements IUserRepository.
func (u *userRepository) FindUserById(userId uint) (*model.User, error) {
	var user model.User
	err := u.db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
