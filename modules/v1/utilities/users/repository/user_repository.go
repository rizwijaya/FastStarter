package repository

import database "FastStarter/app/databases"

type Repository interface {
}

type repository struct {
	db database.Database
}

func NewRepository(db database.Database) *repository {
	return &repository{db}
}

func (r *repository) Hello() string {
	return "Hello, World 👋!"
}
