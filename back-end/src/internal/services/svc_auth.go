package services

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"src/internal/config"
	"src/internal/database"
	"src/internal/models"
	"src/pkg/utils"
)

type ErrorHandlerFunc func(*gin.Context, error, int)

// AuthService is a service that handles the authentication of users
type AuthService struct {
	repo database.AuthRepository
}

// AllUserGuids fetches the list of all the user guids in the DB
func (asv *AuthService) AllUserGuids(c *gin.Context) {
	c.JSON(http.StatusOK, asv.repo.AllUserGuids(c))
}

// Login authenticates a user using Basic Auth
func (asv *AuthService) Login(c *gin.Context) {
	var secrets models.UserSecret
	var err error
	var dbCreds models.UserSecret

	// cast the request into required data structure
	if err = c.ShouldBindJSON(&secrets); err != nil {
		utils.Logger.Error("failed to bind request", zap.Error(err))
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			models.Error{Message: "failed to bind request. Error: " + err.Error()},
		)
		return
	}

	// try basic auth with credentials
	if dbCreds, err = asv.repo.FetchUserCredentials(c, secrets); err != nil {
		utils.Logger.Error("failed to login", zap.Error(err))
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			models.Error{Message: "failed to login. Error: " + err.Error()},
		)
		return
	}

	// check password
	if utils.SimpleHash(secrets.Password) != dbCreds.Password {
		utils.Logger.Error("incorrect username/password")
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			models.Error{Message: "incorrect username/password"},
		)
		return
	}

	// create a JWT token
	token := jwtTokenHandler(c, dbCreds)
	if token == nil {
		return
	}

	// attach appropriate headers
	c.Header("Content-Type", "application/json")
	c.Header("Authorization", "Bearer "+token.BearerToken)
	c.String(http.StatusOK, token.BearerToken)
}

// Register registers a user by storing the userId and hashed password
func (asv *AuthService) Register(c *gin.Context) {
	var secrets models.UserSecret
	var err error

	// cast request into required data structure
	if err = c.ShouldBindJSON(&secrets); err != nil {
		utils.Logger.Error("failed to bind register request", zap.Error(err))
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			models.Error{Message: "failed to bind. Error: " + err.Error()},
		)
		return
	}

	// Store the username and password in DB
	secrets.ID = secrets.Username
	if err = asv.repo.StoreUserCredentials(c, secrets); err != nil {
		utils.Logger.Error("failed to store the user credentials", zap.Error(err))
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			models.Error{Message: "failed to create user/password. Error: " + err.Error()},
		)
		return
	}

	c.String(http.StatusOK, "Created user successfully")
}

// RefreshToken refreshes JWT token assigned to the user
func (asv *AuthService) RefreshToken(c *gin.Context) {
	// fetch the current user from context (token)
	user := currentUserHandler(c)
	if user == nil {
		return
	}
	// create JWT token
	token := jwtTokenHandler(c, models.NewSecretFromUser(user))
	if token == nil {
		return
	}
	// attach appropriate headers
	c.Header("Content-Type", "application/json")
	c.Header("Authorization", "Bearer "+token.BearerToken)
	c.String(http.StatusOK, token.BearerToken)
}

// DeleteUser deletes the user
func (asv *AuthService) DeleteUser(c *gin.Context, userId string) {
	var err error
	var user *models.User

	// fetch the current user from context (token)
	if user = currentUserHandler(c); user == nil {
		return
	}
	// delete user by userId (or username both are same)
	if err = asv.repo.DeleteUser(c, models.UserSecret{ID: user.Username}); err != nil {
		panic("ERR DEL USER " + err.Error())
	}
}

// NewAuthService creates a new instance of AuthService satisfying AuthIface
func NewAuthService(cfg *config.AppConfig) AuthIface {
	repo, err := database.NewAuthRepository(cfg)
	if err != nil {
		utils.Logger.Fatal("failed to get auth repository", zap.Error(err))
		panic(err)
	}
	return &AuthService{repo: repo}
}
