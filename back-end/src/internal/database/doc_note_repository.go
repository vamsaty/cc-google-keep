package database

/*
Package to interact with Document database to create, update or delete GoogleNotes
It uses models package to store and retrieve GoogleNotes.

Author: Satyam Shivam Sundaram
*/

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"src/internal/config"
	"src/internal/models"
	mongoDriver "src/pkg/mongo_driver"
	"src/pkg/utils"
)

// docDBNoteRepository implements methods to be used by services to persist GoogleNotes to a
// document database
type docDBNoteRepository struct {
	mongodbClient *mongoDriver.Client
}

// NOTE: this method doesn't belong to any interface.
func (r *docDBNoteRepository) coll() *mongo.Collection {
	return r.mongodbClient.Collection
}

// UpdateNote updates a note created by the user
func (r *docDBNoteRepository) UpdateNote(c context.Context, note *models.Note) (models.Note, error) {
	result, err := r.coll().UpdateOne(
		c,
		bson.M{"_id": note.ID, "author_id": note.AuthorId},
		bson.M{"$set": note},
	)
	if err != nil {
		utils.Logger.Error("failed to update Note", zap.Error(err))
		return *note, fmt.Errorf("failed to update Note. Error: %s", err.Error())
	}
	if result.ModifiedCount != 1 {
		utils.Logger.Error("Note not modified")
		return *note, fmt.Errorf("must update only 1 Note. Updated: %d notes", result.ModifiedCount)
	}
	return *note, nil
}

// CreateNote creates a Note and attaches it with the author
func (r *docDBNoteRepository) CreateNote(c context.Context, note *models.Note) (models.Note, error) {
	var err error
	utils.Logger.Info("CreateNote", zap.Any("note", *note))

	note.ID = uuid.New().String()
	if _, err = r.coll().InsertOne(c, note); err != nil {
		utils.Logger.Error("failed to create note", zap.Any("note", note))
		return *note, fmt.Errorf("failed to create note. Error: %s", err)
	}
	return *note, nil
}

// GetNoteById returns a note created by the user
func (r *docDBNoteRepository) GetNoteById(c context.Context, noteId string) (models.Note, error) {
	return r.getNoteByFilter(c, bson.M{"_id": noteId})
}

// GetAllNotes fetches all the notes belonging to a user
func (r *docDBNoteRepository) GetAllNotes(c context.Context, userId string) ([]models.Note, error) {
	var notes []models.Note
	cursor, err := r.coll().Find(c, bson.M{"author_id": userId})
	if err != nil {
		return notes, fmt.Errorf("failed to find notes for user %s. Error: %s", userId, err)
	}
	err = cursor.All(c, &notes)
	return notes, err
}

func (r *docDBNoteRepository) getNoteByFilter(c context.Context, filter bson.M) (models.Note, error) {
	var note models.Note
	if err := r.coll().FindOne(c, filter).Decode(&note); err != nil {
		return note, fmt.Errorf("failed to get note. Error: %s", err)
	}
	return note, nil
}

// NewNoteRepository creates a client to interact with mongodb and exposes methods for services
// for persistent storage of GoogleNotes
func NewNoteRepository(cfg *config.AppConfig) (NoteRepository, error) {
	utils.Logger.Info("initializing mongodb client for NoteRepository")
	m, err := initializeMongoDbClient(cfg.DocDBConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to configure mongodb. Error: %s", err)
	}
	return &docDBNoteRepository{mongodbClient: m}, nil
}
