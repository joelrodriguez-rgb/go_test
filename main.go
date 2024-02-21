package main

import (
	"fmt"
	"go_test/model"
	"go_test/repository"
	"go_test/service"
)

func main() {
	var op model.UserOperation = model.NewUserOperation()
	var repository = repository.NewUserRepositoryImpl(op)

	var service = service.NewUserServiceImpl(repository)
	service.CreateUser(model.UserEntity{Id: 1, Name: "Pedro"})

	fmt.Println(repository.Users)

	service.DeleteUser(1)

	fmt.Println(repository.Users)

}
