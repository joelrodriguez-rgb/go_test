package service

import (
	"go_test/model"
	"go_test/repository"
)

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserServiceImpl(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) CreateUser(user model.UserEntity) (model.UserEntity, error) {
	return s.repo.CreateUser(user.Id, user.Name)
}


func (s *UserServiceImpl) DeleteUser(id int) (int, error) {
	return s.repo.DeleteUser(id)
}

