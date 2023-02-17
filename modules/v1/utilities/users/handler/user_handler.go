package handler

import (
	database "FastStarter/app/databases"
	repositoryUsers "FastStarter/modules/v1/utilities/users/repository"
	serviceUsers "FastStarter/modules/v1/utilities/users/service"
)

type usersHandler struct {
	usersService serviceUsers.Service
}

func NewUsersHandler(usersService serviceUsers.Service) *usersHandler {
	return &usersHandler{usersService}
}

func Handler(db database.Database) *usersHandler {
	//Users
	repository := repositoryUsers.NewRepository(db)
	service := serviceUsers.NewService(repository)

	return NewUsersHandler(service)
}
