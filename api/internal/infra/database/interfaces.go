package database

import "github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
