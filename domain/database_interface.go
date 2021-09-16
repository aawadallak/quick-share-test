package domain

type IDatabase interface {
	Store(*Upload) (interface{}, error)
	FindById(id string) (*Database, error)
	FindAll() ([]*Database, error)
	Delete() error
}
