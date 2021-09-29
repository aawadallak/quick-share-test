package domain

import "quick_share/domain"

type IDatabase interface {
	Store(*domain.Upload) (interface{}, error)
	FindById(id string) (*domain.Database, error)
	FindAll() ([]*domain.Database, error)
	Delete() error
}
