package services

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"net/http"
	"src/internal/models"
	"src/pkg/utils"
)

func ErrorMessage(msg string) models.Error {
	return models.Error{Message: msg}
}

func ErrorMessageWithErr(msg string, err error) models.Error {
	return models.Error{Message: msg + ". Error: " + err.Error()}
}

func handleAuthError(c *gin.Context, err models.Error) {
	c.Header("WWW-Authenticate", "Bearer")
	c.AbortWithStatusJSON(http.StatusUnauthorized, err)
}

// CurrentUser extracts the current username from JWT token in Auth header
func CurrentUser(ctx *gin.Context) (models.User, error) {
	var user models.User
	utils.Logger.Info("extracting current user")
	tok, err := models.ExtractAuthToken(ctx)
	if err != nil {
		utils.Logger.Info("failed to extract auth token")
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			ErrorMessageWithErr("", err),
		)
		return user, err
	}

	jwtTok, err := models.ParseJWTToken(tok)
	if err != nil {
		utils.Logger.Info("failed to parse KWT token")
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			ErrorMessageWithErr("", err),
		)
		return user, err
	}

	// current username
	claims := jwtTok.Claims.(jwt.MapClaims)
	user = models.User{
		ID:       claims["username"].(string),
		Username: claims["username"].(string),
	}
	utils.Logger.Info("current-User-found-here")
	return user, nil
}

func currentUserHandler(c *gin.Context) *models.User {
	// fetch the current user from context (token)
	var user models.User
	var err error
	if user, err = CurrentUser(c); err != nil {
		utils.Logger.Error("failed to extract user from token", zap.Error(err))
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			ErrorMessageWithErr("failed to extract user from token", err),
		)
		return nil
	}
	return &user
}

func jwtTokenHandler(c *gin.Context, dbCreds models.UserSecret) *models.AuthJwtToken {
	token, err := models.NewAuthJwtToken(dbCreds)
	if err != nil {
		utils.Logger.Error("failed to generate token", zap.Error(err))
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			ErrorMessageWithErr("failed to generate token", err),
		)
		return nil
	}
	return &token
}
