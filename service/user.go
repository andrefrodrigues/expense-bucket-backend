package service

import (
	"context"
	"errors"
	"expense-bucket-api/model"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

type UserServiceConfig struct {
	DB *gorm.DB
}

func NewUserService(config UserServiceConfig) *UserService {
	return &UserService{
		db: config.DB,
	}
}

//GetUserByEmail gets a user by its email
func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{Email: email}
	result := u.db.WithContext(ctx).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, UserNotFoundErr
		}
		return nil, result.Error
	}

	return &user, nil
}
