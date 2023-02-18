package controllers

import (
	database "FastStarter/infrastructures/databases"
	logger "FastStarter/infrastructures/logger"
	repositoryUsers "FastStarter/modules/v1/users/interfaces/repositories"
	usersUsecase "FastStarter/modules/v1/users/usecases"
)

type UsersController struct {
	usersUsecase usersUsecase.UsecasePresenter
	log          logger.LoggerPresenter
}

func NewController(usersUsecase usersUsecase.UsecasePresenter, log logger.LoggerPresenter) *UsersController {
	return &UsersController{usersUsecase, log}
}

func UserController(db database.Database) *UsersController {
	//Users
	log := logger.NewLoggger()
	repository := repositoryUsers.NewRepository(db)
	usecase := usersUsecase.NewUsecase(repository)

	return NewController(usecase, log)
}
