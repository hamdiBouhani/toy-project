package mongo

import (
	"context"
	"errors"
	"toy-project/storage"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// conn is the main database connection.
type conn struct {
	db *mongo.Client
}

func (c *conn) Close() error {
	return c.db.Disconnect(context.TODO())
}

func (c *conn) Version() (i string, err error) {
	return i, errors.New("implement me")
}

func Open(ctx context.Context, dsn string, logger *logrus.Logger) (storage.Storage, error) {
	clientOptions := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatal(err)
	}

	return &conn{
		db: client,
	}, nil
}
