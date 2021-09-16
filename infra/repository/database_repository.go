package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"we/database"
	"we/domain"
	"we/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DatabaseRepository struct{}

func (d *DatabaseRepository) Store(file *domain.Upload) (interface{}, error) {

	conn, err := database.Conn()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	db := conn.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))

	now := time.Now().Add(time.Hour * 24).Local().Unix()

	field := struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		Path      string             `bson:"path,omitempty"`
		ExpiresAt int64              `bson:"expires_at,omitempty"`
	}{
		Path:      file.GetFolder().GetPath(),
		ExpiresAt: now,
	}

	res, err := db.InsertOne(context.TODO(), field)

	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (d *DatabaseRepository) FindById(id string) (*domain.Database, error) {

	// convert id string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	conn, err := database.Conn()

	if err != nil {
		return nil, err
	}

	db := conn.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))

	var dto dto.DatabaseDTO

	err = db.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&dto)

	if err != nil {
		return nil, fmt.Errorf("cannot find the informed document")
	}

	dmn := domain.NewDomainDatabase(dto.ID.Hex(), dto.Path)

	return dmn, nil
}

func (d *DatabaseRepository) FindAll() ([]*domain.Database, error) {

	conn, err := database.Conn()

	if err != nil {
		return nil, err
	}

	db := conn.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))

	res, err := db.Find(context.Background(), bson.M{"expires_at": bson.M{"$lte": time.Now().Local().Unix()}})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Close(context.Background())

	var array []*domain.Database

	for res.Next(context.Background()) {

		var dt *dto.DatabaseDTO

		err = res.Decode(&dt)

		array = append(array, dt.Convert2Entity())

		if err != nil {
			return nil, err
		}
	}

	return array, nil
}

func (d *DatabaseRepository) Delete() error {

	conn, err := database.Conn()

	if err != nil {
		return err
	}

	db := conn.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))

	_, err = db.DeleteMany(context.Background(), bson.M{"expires_at": bson.M{"$lte": time.Now().Local().Unix()}})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
