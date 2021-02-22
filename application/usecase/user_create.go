package usecase

import "go-clean-architecture-sample/application/usecase/inputport"

type UserCreate interface {
	Handle(data *inputport.UserCreateData) error
}
