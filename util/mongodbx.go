package util

import (
	"context"
	"log"

	"github.com/StevenZack/databaser/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func QueryMongoDB(conn data.Connection, query string) (*data.Result, error) {
	c, e := mongo.Connect(context.TODO(), options.Client().ApplyURI(conn.Dsn))
	if e != nil {
		log.Println(e)
		return nil, e
	}
	defer c.Disconnect(context.TODO())
	// c.Database("name").Collection("name").
	return nil, nil
}
