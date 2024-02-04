package services

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"net/http"
	"src/internal/models"
	"src/pkg/utils"
)

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
			models.Error{
				Code:    models.TokenParseFailed,
				Message: err.Error(),
			},
		)
		return user, err
	}

	jwtTok, err := models.ParseJWTToken(tok)
	if err != nil {
		utils.Logger.Info("failed to parse KWT token")
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			models.Error{
				Code:    models.TokenParseFailed,
				Message: err.Error(),
			},
		)
		return user, err
	}
	// current username
	claims := jwtTok.Claims.(jwt.MapClaims)
	user = models.User{
		ID:       claims["username"].(string),
		Username: claims["username"].(string),
	}
	return user, nil
}

func currentUserHandler(c *gin.Context) *models.User {
	// fetch the current user from context (token)
	if user, err := CurrentUser(c); err != nil {
		utils.Logger.Error("failed to extract user from token", zap.Error(err))
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			models.Error{Message: "failed to extract user from token. Error: " + err.Error()},
		)
		return &user
	}
	return nil
}

func jwtTokenHandler(c *gin.Context, dbCreds models.UserSecret) *models.AuthJwtToken {
	token, err := models.NewAuthJwtToken(dbCreds)
	if err != nil {
		utils.Logger.Error("failed to generate token", zap.Error(err))
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			models.Error{Message: "failed to generate token. Error: " + err.Error()},
		)
		return nil
	}
	return &token
}
