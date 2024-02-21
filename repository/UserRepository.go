package repository

import (
	"go_test/model"
)

type UserRepository interface {
	CreateUser(id int, name string) (model.UserEntity, error)
	DeleteUser(user *model.UserEntity) (int, error)
}

type UserRepositoryImpl struct {
	operation model.UserOperation
}

func NewUserRepositoryImpl(op model.UserOperation) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		operation: op,
	}
}

func (e *UserRepositoryImpl) CreateUser(id int, name string) (model.UserEntity, error) {
	u , _ := e.operation.Create(id, name)
	return u, nil
}

func (e *UserRepositoryImpl) DeleteUser(user *model.UserEntity) (int, error) {
	//eliminar el valor que esta en el puntero
	user = nil
	return 1, nil
}
