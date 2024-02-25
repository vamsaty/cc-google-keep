package database

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"src/internal/config"
	"src/internal/models"
	mongoDriver "src/pkg/mongo_driver"
	"src/pkg/utils"
)

// NoteRepository is an interface to be implemented by any repository that interacts with the
// document database to create, update or delete notes
type NoteRepository interface {
	UpdateNote(c context.Context, note *models.Note) (models.Note, error)
	CreateNote(c context.Context, note *models.Note) (models.Note, error)
	GetNoteById(c context.Context, noteId string) (models.Note, error)
	GetAllNotes(c context.Context, userId string) ([]models.Note, error)
}

// AuthRepository is an interface to be implemented by any repository that interacts with the
// document database to create, update or delete user credentials.
// This is supposed to be a session store
type AuthRepository interface {
	CreateUserAccount(context.Context, models.UserSecret) error
	FetchUserCredentials(context.Context, models.UserSecret) (models.UserSecret, error)
	DeleteUser(context.Context, models.UserSecret) error
	AllUserGuids(context.Context) []string
}

type UserRepository interface {
	GetInfo(context.Context, string) (models.User, error)
	CreateUserProfile(context.Context, string) (models.User, error)
	UpdateLastLogin(context.Context, string) error
}

// initializeMongoDbClient creates a mongodb client using @config.AppConfig
func initializeMongoDbClient(cfg config.MongoDBConfig) (*mongoDriver.Client, error) {
	utils.Logger.Info("initializing mongodb client")
	credentials := readMongoCredentials(cfg.SecretsPath)

	utils.Logger.Info("creating mongodb client")
	client, err := mongoDriver.New(
		mongoDriver.WithMongoDBURL(cfg.Url),
		mongoDriver.WithCollectionName(cfg.CollectionName),
		mongoDriver.WithDatabaseName(cfg.DatabaseName),
		mongoDriver.WithPassword(credentials["password"].(string)),
		mongoDriver.WithUsername(credentials["username"].(string)),
	)
	if err != nil {
		return nil, fmt.Errorf("mongodb client creation failed. Error: %s", err)
	}
	return client, err
}

// readMongoCredentials reads the secret file (as json) and returns a map containing secrets
func readMongoCredentials(secretFile string) (secrets map[string]interface{}) {
	var file io.ReadCloser
	var err error
	if file, err = os.Open(secretFile); err != nil {
		return nil
	}
	defer file.Close()

	_ = json.NewDecoder(file).Decode(&secrets)
	return secrets
}
