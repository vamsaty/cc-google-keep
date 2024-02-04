package mongo_driver

import (
	"go.uber.org/zap"
)

type Option func(c *Client)

func WithLogger(logger zap.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

func WithMongoDBURL(url string) Option {
	return func(c *Client) {
		c.dbURL = url
	}
}

func WithUsername(u string) Option {
	return func(c *Client) {
		c.username = u
	}
}

func WithPassword(p string) Option {
	return func(c *Client) {
		c.password = p
	}
}

func WithDatabaseName(d string) Option {
	return func(c *Client) {
		c.databaseName = d
	}
}

func WithCollectionName(collection string) Option {
	return func(c *Client) {
		c.collectionName = collection
	}
}
