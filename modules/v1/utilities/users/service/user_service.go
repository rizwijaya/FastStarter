package service

import repositoryUsers "FastStarter/modules/v1/utilities/users/repository"

type Service interface {
}

type service struct {
	repository repositoryUsers.Repository
}

func NewService(repository repositoryUsers.Repository) *service {
	return &service{repository}
}
