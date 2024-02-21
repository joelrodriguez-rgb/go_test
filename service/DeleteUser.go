package service

import (
	"go_test/model"
)

type DeleteUserService interface {
	DeleteUser(user *model.UserEntity) (int, error)
}

