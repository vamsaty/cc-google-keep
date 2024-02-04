package mongo_driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoClientIface interface {
	GetCollection()
	Connect(ctx context.Context)
}

type Client struct {
	dbURL          string
	username       string
	password       string
	databaseName   string
	collectionName string
	logger         zap.Logger
	Collection     *mongo.Collection
}

// Connect is used to establish a connection to the mongodb database server
func (c *Client) Connect(ctx context.Context) error {
	credentials := options.Credential{
		Username: c.username,
		Password: c.password,
	}
	opts := options.Client().ApplyURI(c.dbURL).SetAuth(credentials).
		SetRetryWrites(false)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return fmt.Errorf("failed to connect to mongodb. Error: %s", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping mongodb. Error: %s", err)
	}
	c.Collection = client.Database(c.databaseName).Collection(c.collectionName)
	return nil
}

// New creates a new mongodb database client using @opts
func New(opts ...Option) (*Client, error) {
	c := &Client{}
	for _, o := range opts {
		o(c)
	}
	err := c.Connect(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error in setting up connection %w", err)
	}
	return c, err
}
