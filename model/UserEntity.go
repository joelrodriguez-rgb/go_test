package model

type UserOperation interface {
	create(id int, name string) (UserEntity, error)
}

type UserEntity struct {
	Id   int
	Name string
}

func (user *UserEntity) create(id int, name string) (UserEntity, error) {
	return UserEntity{Id: id, Name: name}, nil
}
