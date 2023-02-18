package repository

import database "FastStarter/infrastructures/databases"

type RepositoryPresenter interface {
}

type Repository struct {
	db database.Database
}

func NewRepository(db database.Database) *Repository {
	return &Repository{db}
}
