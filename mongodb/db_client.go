package mongodb

import (
	"context"
	"goms/mongodb/usecase"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type DBClient struct {
	dbClient              *mongo.Client
	timeout               time.Duration
	CounterUsecase        usecase.CounterUsecase
	AccountUsecase        usecase.AccountUsecase
	CharacterOrderUsecase usecase.CharacterOrderUsecase
	CharacterUsecase      usecase.CharacterUsecase
	InventoryUsecase      usecase.InventoryUsecase
	SocialUsecase         usecase.SocialUsecase
}

// Calling Connect does not block for server discovery.
// If you wish to know if a MongoDB server has been found and connected to, use the Ping method
func NewDBClient(uri string, name string) *DBClient {
	c := &DBClient{timeout: 3 * time.Second}
	ctx, canel := context.WithTimeout(context.Background(), c.timeout)
	defer canel()
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		// Just catch the error when init a new MongoDB client failed
		// The mongo.Connect function does not immediately attempt to connect to the mongo
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	err = dbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	c.dbClient = dbClient
	slog.Info("MongoDB conn was established")
	db := c.dbClient.Database(name)
	c.CounterUsecase = usecase.NewCounterUsecase(db)
	c.AccountUsecase = usecase.NewAccountUsecase(db)
	c.CharacterOrderUsecase = usecase.NewCharacterOrderUsecase(db)
	c.CharacterUsecase = usecase.NewCharacterUsecase(db)
	c.InventoryUsecase = usecase.NewInventoryUsecase(db)
	c.SocialUsecase = usecase.NewSocialUsecase(db)
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

func (c *DBClient) WithTransaction(handleDoc func(ctx mongo.SessionContext) (any, error)) {
	wc := writeconcern.Majority()
	txnOptions := options.Transaction().SetWriteConcern(wc)
	// Starts a session on the client
	session, err := c.dbClient.StartSession()
	if err != nil {
		slog.Error("Failed to start session", "err", err)
		return
	}
	defer session.EndSession(context.TODO())
	_, err = session.WithTransaction(context.TODO(), handleDoc, txnOptions)
	if err != nil {
		slog.Error("Failed to handle doc with transaction", "err", err)
		return
	}
	slog.Debug("Handle doc with transaction ok")
}
