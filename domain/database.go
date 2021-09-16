package domain

type Database struct {
	id   string
	path string
}

func (d *Database) GetId() string {
	return d.id
}

func (d *Database) GetPath() string {
	return d.path
}

func NewDomainDatabase(id, path string) *Database {
	return &Database{
		id:   id,
		path: path,
	}
}
