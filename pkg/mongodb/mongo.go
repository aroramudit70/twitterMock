package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Uri          string `envconfig:"DB_URI"`
	DataBaseName string `envconfig:"DB_NANE"`
}

func GetConnection(m Config) (*mongo.Database, error) {
	connection, err := mongo.NewClient(options.Client().ApplyURI(m.Uri))
	if err != nil {
		return nil, err
	}
	err = connection.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	db := connection.Database(m.DataBaseName)
	return db, nil
}
