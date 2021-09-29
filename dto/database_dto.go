package dto

import (
	"quick_share/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatabaseDTO struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Path string             `bson:"path,omitempty"`
}

func (d *DatabaseDTO) Convert2Entity() *domain.Database {
	return domain.NewDomainDatabase(d.ID.Hex(), d.Path)
}
