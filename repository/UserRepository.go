package repository

import (
	"go_test/model"
)

type UserRepository interface {
	CreateUser(id int, name string) (model.UserEntity, error)
	DeleteUser(id int) (int, error)
}

type UserRepositoryImpl struct {
	Users    []model.UserEntity
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{
	}
}

func (e *UserRepositoryImpl) CreateUser(id int, name string) (model.UserEntity, error) {
	u := model.UserEntity{Id: id, Name: name}
	e.Users = append(e.Users, u)
	return u, nil
}

func (e *UserRepositoryImpl) DeleteUser(id int) (int, error) {
	//eliminar el usuario segun el id de la lista de users
	for i, user := range e.Users {
		if user.Id == id {
			e.Users = append(e.Users[:i], e.Users[i+1:]...)
			return id, nil
		}
	}
	return 0, nil
}
