package di

import (
	"go-clean-architecture-sample/adapter/controller"
	"go-clean-architecture-sample/adapter/presenter"
	"go-clean-architecture-sample/adapter/repository"
	"go-clean-architecture-sample/application/usecase"
	"go-clean-architecture-sample/application/usecase/interactor"
	"go-clean-architecture-sample/driver/database/dummy"
)

type Container interface {
	NewUserRepository() repository.User
	NewUserCreatePresenter() presenter.UserCreate
	NewUserCreateUseCase(userRepository repository.User, presenter presenter.UserCreate) usecase.UserCreate
	NewUserController(userCreateUseCase usecase.UserCreate) *controller.User
}

var _ Container = &container{}

func NewContainer() *container {
	c := &container{}
	c.userRepository = c.NewUserRepository()
	c.userCreatePresenter = c.NewUserCreatePresenter()
	c.userCreateUseCase = c.NewUserCreateUseCase(c.userRepository, c.userCreatePresenter)
	c.UserController = c.NewUserController(c.userCreateUseCase)

	return c
}

type container struct {
	userRepository      repository.User
	userCreatePresenter presenter.UserCreate
	userCreateUseCase   usecase.UserCreate
	UserController      *controller.User
}

func (d *container) NewUserRepository() repository.User {
	return dummy.NewUser()
}

func (d *container) NewUserCreatePresenter() presenter.UserCreate {
	return presenter.NewUserCreate()
}

func (d *container) NewUserCreateUseCase(userRepository repository.User, presenter presenter.UserCreate) usecase.UserCreate {
	return interactor.NewUserCreate(userRepository, presenter)
}

func (d *container) NewUserController(userCreateUseCase usecase.UserCreate) *controller.User {
	return controller.NewUser(userCreateUseCase)
}
