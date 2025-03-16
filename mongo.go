package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// Config holds the configuration settings for connecting to MongoDB.
type Config struct {
	URI      string `yaml:"uri"`
	Database string `yaml:"database"`
}

// Client is an alias to mongo.Database.
type Client = *mongo.Database

// New creates a new connection to the MongoDB database.
func New(c context.Context, cfg *Config) (Client, error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	db := client.Database(cfg.Database)

	return db, nil
}
