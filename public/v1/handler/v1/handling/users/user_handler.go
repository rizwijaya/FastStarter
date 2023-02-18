package handler

import (
	database "FastStarter/infrastructures/databases"
	repositoryUsers "FastStarter/modules/v1/users/interfaces/repositories"
	usecaseUsers "FastStarter/modules/v1/users/usecases"
)

type usersHandler struct {
	usersService usecaseUsers.UsecasePresenter
}

func NewUsersHandler(usersService usecaseUsers.UsecasePresenter) *usersHandler {
	return &usersHandler{usersService}
}

func Handler(db database.Database) *usersHandler {
	//Users
	repository := repositoryUsers.NewRepository(db)
	service := usecaseUsers.NewUsecase(repository)

	return NewUsersHandler(service)
}
