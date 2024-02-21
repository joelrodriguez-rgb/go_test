package model

type UserOperation interface {
	Create(id int, name string) (UserEntity, error)
}

func NewUserOperation() UserOperation {
	return &UserEntity{}
}

type UserEntity struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}


func (user *UserEntity) Create(id int, name string) (UserEntity, error) {
	return UserEntity{Id: id, Name: name}, nil
}
