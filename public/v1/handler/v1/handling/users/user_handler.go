package handler

import (
	database "FastStarter/infrastructures/databases"
	"FastStarter/infrastructures/logger"
	repositoryUsers "FastStarter/modules/v1/users/interfaces/repositories"
	usecaseUsers "FastStarter/modules/v1/users/usecases"
)

type usersHandler struct {
	usersUsecase usecaseUsers.UsecasePresenter
	log          logger.LoggerPresenter
}

func NewUsersHandler(usersUsecase usecaseUsers.UsecasePresenter, log logger.LoggerPresenter) *usersHandler {
	return &usersHandler{usersUsecase, log}
}

func Handler(db database.Database) *usersHandler {
	//Users
	log := logger.NewLoggger()
	repository := repositoryUsers.NewRepository(db)
	usecase := usecaseUsers.NewUsecase(repository)

	return NewUsersHandler(usecase, log)
}
