package controller

import (
	"go-clean-architecture-sample/application/usecase"
	"go-clean-architecture-sample/application/usecase/inputport"
)

func NewUser(userCreateUserCase usecase.UserCreate) *User {
	return &User{
		userCreateUserCase: userCreateUserCase,
	}
}

type User struct {
	userCreateUserCase usecase.UserCreate
}

func (c *User) CreateUser(userName string) error {
	input := inputport.NewUserCreateData(userName)
	if err := c.userCreateUserCase.Handle(input); err != nil {
		return err
	}
	return nil
}
