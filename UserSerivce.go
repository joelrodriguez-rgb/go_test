package main

import (
	"go_test/model"
)

type Service interface {
	CreateUser(user model.UserEntity) (model.UserEntity, error)
}

type UserService struct {
	service UserServiceImpl
}

func NewUserService(service UserServiceImpl) *UserService {
	return &UserService{service: service}
}
