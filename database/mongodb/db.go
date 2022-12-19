package mongodb

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/PT-Jojonomic-Indonesia/microkit/sentry"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type AuthConfig options.Credential

var DB *mongo.Database
var Client *mongo.Client

var Init = func(host, name, port string, authConfig AuthConfig) {
	ctx := context.Background()

	uri := fmt.Sprintf("mongodb://%s:%s/%s", host, port, name)

	log.Printf("[info] connecting to mongodb with uri : %s", uri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(options.Credential(authConfig)))
	if err != nil {
		sentry.CaptureError(err)
		log.Panic(err)
	}

	db := client.Database(name)
	Client = client
	log.Println("[info] connected to mongodb", db.Name())
}

var GetCollection = func(collection string) *mongo.Collection {
	return DB.Collection(collection)
}

var Health = func() error {
	ctx := context.Background()

	if err := Client.Ping(ctx, readpref.Primary()); err != nil {
		sentry.CaptureError(err)
		return errors.New("mongo db is not available")
	}

	return nil
}
