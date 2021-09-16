package database

import (
	"context"
	"fmt"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	driver   string
	user     string
	pass     string
	host     string
	port     string
	database string
}

func createDatabase() Database {
	return Database{
		driver:   os.Getenv("DB_TYPE"),
		user:     os.Getenv("DB_USER"),
		pass:     os.Getenv("DB_PASS"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		database: os.Getenv("DB_NAME"),
	}
}

var db *mongo.Client
var mongoOnce sync.Once
var errorHandler error

func Conn() (*mongo.Client, error) {

	mongoOnce.Do(func() {

		_db := createDatabase()

		dsn := fmt.Sprintf("%s://%s:%s@%s:%s", _db.driver, _db.user, _db.pass, _db.host, _db.port)

		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))

		if err != nil {
			errorHandler = err
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			errorHandler = err
		}

		db = client

	})

	return db, errorHandler
}
