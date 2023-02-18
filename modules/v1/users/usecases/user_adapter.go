package service

import repositoriesUsers "FastStarter/modules/v1/users/interfaces/repositories"

type UsecasePresenter interface {
}

type Usecase struct {
	repository repositoriesUsers.RepositoryPresenter
}

func NewUsecase(repositories repositoriesUsers.RepositoryPresenter) *Usecase {
	return &Usecase{repositories}
}
