package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(uri string) (*mongo.Client, error) {
    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if (err != nil) {
        return nil, err
    }

    err = client.Ping(context.TODO(), nil)
    return client, err
}