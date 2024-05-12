package mongodb

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBClient struct {
	dbClient *mongo.Client
	timeout  time.Duration
}

// Calling Connect does not block for server discovery.
// If you wish to know if a MongoDB server has been found and connected to, use the Ping method
func NewDBClient(uri string, timeout uint8) *DBClient {
	c := &DBClient{timeout: time.Duration(timeout) * time.Second}
	ctx, canel := context.WithTimeout(context.Background(), c.timeout)
	defer canel()
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		// Just catch the error when init a new MongoDB client failed
		// The mongo.Connect function does not immediately attempt to connect to the mongo
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = dbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	c.dbClient = dbClient
	slog.Info("MongoDB conn was established")
	return c
}

func (c *DBClient) Disconnect() {
	ctx, canel := context.WithTimeout(context.Background(), c.timeout)
	defer canel()
	err := c.dbClient.Disconnect(ctx)
	if err != nil {
		slog.Error("Failed to disconnect DBClient", "err", err)
		return
	}
	slog.Info("DBClient was disconnected")
	c.dbClient = nil
}

func (c *DBClient) SelectDB(name string) *mongo.Database {
	return c.dbClient.Database(name)
}
