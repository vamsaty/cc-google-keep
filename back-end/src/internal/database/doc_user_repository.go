package database

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
	"time"
)

type docDBUserRepository struct {
	mongoDbClient *mongoDriver.Client
}

func (dUser *docDBUserRepository) UpdateLastLogin(ctx context.Context, userId string) error {
	var user models.User
	var err error

	if err = dUser.coll().FindOne(ctx, bson.M{"_id": userId}).Decode(&user); err != nil {
		_, err = dUser.CreateUserProfile(ctx, userId)
		return err
	}

	// update last login
	user.LastLoginAt = time.Now()
	if _, err = dUser.coll().UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
	); err != nil {
		utils.Logger.Error("last login update failed. Failed to update database", zap.Error(err))
	}
	return err
}

func (dUser *docDBUserRepository) CreateUserProfile(ctx context.Context, userId string) (models.User, error) {
	var user models.User

	_, err := dUser.coll().Find(ctx, bson.M{"_id": userId})
	// user profile was found, just update
	if err != nil {
		err = fmt.Errorf("user profile already exists")
		utils.Logger.Error("user already exists", zap.Error(err))
		return user, err
	}

	// create the empty user profile
	user.ID = userId
	user.Username = userId
	user.LastLoginAt = time.Now()
	if _, err = dUser.coll().InsertOne(ctx, &user); err != nil {
		utils.Logger.Error("failed to persist user profile", zap.Error(err))
	}
	return user, nil
}

func (dUser *docDBUserRepository) coll() *mongo.Collection {
	return dUser.mongoDbClient.Collection
}

func (dUser *docDBUserRepository) GetInfo(ctx context.Context, userId string) (models.User, error) {
	var user models.User
	utils.Logger.Info("getting user info", zap.String("user_id", userId))

	if err := dUser.coll().FindOne(ctx, bson.M{"_id": userId}).Decode(&user); err != nil {
		utils.Logger.Error("failed to fetch user info", zap.Error(err))
		return user, err
	}
	return user, nil
}

func NewUserRepository(cfg *config.AppConfig) (UserRepository, error) {
	m, err := initializeMongoDbClient(cfg.UserDBConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to configure mongodb client. Error: %s", err)
	}
	return &docDBUserRepository{
		mongoDbClient: m,
	}, nil
}
