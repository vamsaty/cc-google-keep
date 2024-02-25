package database

/*
This package interacts with the document database to create, update or delete user credentials.
It uses models package to store and retrieve user credentials.
Author: Satyam Shivam Sundaram
*/

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"src/internal/config"
	"src/internal/models"
	mongoDriver "src/pkg/mongo_driver"
	"src/pkg/utils"
)

type docDBAuthRepository struct {
	mongoDbClient *mongoDriver.Client
}

func (ar *docDBAuthRepository) coll() *mongo.Collection {
	return ar.mongoDbClient.Collection
}

// CreateUserAccount stores the user credentials to DB
func (ar *docDBAuthRepository) CreateUserAccount(ctx context.Context, creds models.UserSecret) error {
	var err error
	var secret models.UserSecret

	// check if a user already exists
	utils.Logger.Info("creating user", zap.String("username", creds.Username))
	err = ar.coll().FindOne(ctx, bson.M{"username": creds.Username}).Decode(&secret)
	if err == nil {
		utils.Logger.Error("failed to create user. user already exists")
		return fmt.Errorf("user already exists, please login")
	}

	// update username and password
	pwd := utils.SimpleHash(creds.Password)
	_, err = ar.coll().InsertOne(ctx, models.NewUserSecret(creds.Username, pwd))

	return err
}

// FetchUserCredentials authenticates a user and shares an AuthToken
func (ar *docDBAuthRepository) FetchUserCredentials(
	ctx context.Context, creds models.UserSecret) (models.UserSecret, error) {
	var secret models.UserSecret
	var err error

	// find user, match the password
	err = ar.coll().FindOne(ctx, bson.M{"username": creds.Username}).Decode(&secret)
	if err != nil {
		utils.Logger.Error("no user found")
		return secret, fmt.Errorf("user not found, please create an account")
	}
	return secret, nil
}

// DeleteUser deletes a user from DB
func (ar *docDBAuthRepository) DeleteUser(ctx context.Context, creds models.UserSecret) error {
	if result, err := ar.coll().DeleteOne(ctx, bson.M{"_id": creds.ID}); err != nil {
		utils.Logger.Error("failed to delete user", zap.Error(err))
		return err
	} else if result.DeletedCount != 1 {
		errMessage := fmt.Sprintf("deleted %d entries, expected 1 entry", result.DeletedCount)
		utils.Logger.Error(errMessage, zap.Error(err))
		return fmt.Errorf(errMessage)
	}
	utils.Logger.Info("user deleted", zap.String("username", creds.Username))
	return nil
}

func (ar *docDBAuthRepository) AllUserGuids(ctx context.Context) []string {
	var users []models.UserSecret
	if res, err := ar.coll().Find(ctx, bson.M{}); err != nil {
		return nil
	} else if err = res.All(ctx, &users); err != nil {
		return nil
	}
	var data []string
	for _, user := range users {
		data = append(data, user.ID)
	}
	return data
}

func NewAuthRepository(cfg *config.AppConfig) (AuthRepository, error) {
	m, err := initializeMongoDbClient(cfg.AuthDBConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to configure mongodb client. Error: %s", err)
	}
	return &docDBAuthRepository{
		mongoDbClient: m,
	}, nil
}
