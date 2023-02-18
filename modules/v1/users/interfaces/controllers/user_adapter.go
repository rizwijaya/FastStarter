package controllers

import (
	database "FastStarter/infrastructures/databases"
	repositoryUsers "FastStarter/modules/v1/users/interfaces/repositories"
	usersUsecase "FastStarter/modules/v1/users/usecases"
)

type UsersController struct {
	usersUsecase usersUsecase.UsecasePresenter
}

func NewController(usersUsecase usersUsecase.UsecasePresenter) *UsersController {
	return &UsersController{usersUsecase}
}

func UserController(db database.Database) *UsersController {
	//Users
	repository := repositoryUsers.NewRepository(db)
	usecase := usersUsecase.NewUsecase(repository)

	return NewController(usecase)
}
