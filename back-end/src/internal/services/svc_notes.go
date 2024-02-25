package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"src/internal/config"
	"src/internal/database"
	"src/internal/models"
	"src/pkg/utils"
	"time"
)

// GNoteService takes the request and moulds it into Note and other consumable format
type GNoteService struct {
	cfg  *config.AppConfig
	Repo database.NoteRepository
}

// GetNote fetches the note from the database for the current user
func (gn *GNoteService) GetNote(ctx *gin.Context) {
	noteId := ctx.Param("noteId")
	note, err := gn.Repo.GetNoteById(ctx, noteId)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			ErrorMessageWithErr("", err),
		)
	}
	ctx.JSON(http.StatusOK, note)
}

// CreateNote creates a new note in the database for the current user
func (gn *GNoteService) CreateNote(ctx *gin.Context) {
	// get the current user
	var err error
	var user models.User

	if user, err = CurrentUser(ctx); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			ErrorMessageWithErr("", err),
		)
		return
	}
	now := time.Now()
	requestNote := models.Note{
		AuthorId:     user.ID,
		CreatedAt:    now,
		LastModified: now,
	}
	utils.Logger.Info("creating note")
	if err = ctx.ShouldBindJSON(&requestNote); err != nil {
		utils.Logger.Error("failed to bind request Note", zap.Error(err))
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			ErrorMessage("failed to bind create note API input"),
		)
		return
	}

	// converted for internal processing
	if !models.IsValidCreateNote(&requestNote) {
		utils.Logger.Error("Invalid request Note provided", zap.Error(err))
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			ErrorMessage("Invalid payload received"),
		)
		return
	}

	// persist the Note
	var responseNote models.Note
	if responseNote, err = gn.Repo.CreateNote(ctx, &requestNote); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		utils.Logger.Error("failed to persist Note", zap.Error(err))
		return
	}
	utils.Logger.Info("persisted Note", zap.String("Note", fmt.Sprintf("%#v", responseNote)))

	ctx.JSON(http.StatusCreated, responseNote)
	// return the note in response
}

// UpdateNote updates the note in the database for the current user
func (gn *GNoteService) UpdateNote(ctx *gin.Context) {
	var requestNote, dbNote, newNote models.Note
	var err error

	// decode the note in request
	if err = ctx.ShouldBindJSON(&requestNote); err != nil {
		utils.Logger.Error("failed to bind request Note", zap.Error(err))
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			ErrorMessage("failed to bind update note API input"),
		)
		return
	}
	// get note from db to update
	noteId := ctx.Param("noteId")
	if dbNote, err = gn.Repo.GetNoteById(ctx, noteId); err != nil {
		utils.Logger.Error("failed to get Note", zap.Error(err))
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			ErrorMessage("failed to fetch note to update"),
		)
		return
	}
	// update the Note from database
	dbNote.Title = requestNote.Title
	dbNote.Contents = requestNote.Contents
	dbNote.LastModified = time.Now()
	if newNote, err = gn.Repo.UpdateNote(ctx, &dbNote); err != nil {
		utils.Logger.Error("failed to update Note")
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			ErrorMessageWithErr("failed to update Note", err),
		)
		return
	}
	ctx.JSON(http.StatusOK, newNote)
}

// GetAllNotes fetches all the notes from the database for the current user
func (gn *GNoteService) GetAllNotes(ctx *gin.Context) {
	var user models.User
	var err error

	utils.Logger.Info("extracting the current user")
	if user, err = CurrentUser(ctx); err != nil {
		return
	}

	notes, err := gn.Repo.GetAllNotes(ctx, user.ID)
	if err != nil {
		utils.Logger.Info("failed to fetch all notes")
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			ErrorMessageWithErr("", err),
		)
	}
	ctx.JSON(http.StatusOK, &notes)
}

func NewGNoteService(cfg *config.AppConfig) NoteIface {
	repo, _ := database.NewNoteRepository(cfg)
	fmt.Println("note service")
	return &GNoteService{
		cfg:  cfg,
		Repo: repo,
	}
}
