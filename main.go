package main

import (
	"fmt"
	"go_test/model"
	"go_test/repository"
	"go_test/service"
)

func main() {
	var repository = repository.NewUserRepositoryImpl()

	var service = service.NewUserServiceImpl(repository)
	service.CreateUser(model.UserEntity{Id: 1, Name: "Pedro"})
	service.CreateUser(model.UserEntity{Id: 2, Name: "Juan"})

	fmt.Println(repository.Users)

	service.DeleteUser(1)

	fmt.Println(repository.Users)

}
