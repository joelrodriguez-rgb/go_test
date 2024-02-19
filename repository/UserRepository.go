package repository

import "go_test/model"

type UserRepository interface {
	CreateUser(id int, name string) (model.UserEntity, error)
}

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (e *UserRepositoryImpl) CreateUser(id int, name string) (model.UserEntity, error) {
	user := model.UserEntity{Id: 1, Name: "asd"}
	return user, nil
}
