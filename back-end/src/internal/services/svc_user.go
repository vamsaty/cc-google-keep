package services

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"src/internal/config"
	"src/internal/database"
	"src/pkg/utils"
)

type UserService struct {
	repo database.UserRepository
}

func (us *UserService) Info(c *gin.Context) {
	user := currentUserHandler(c)
	userInfo, err := us.repo.GetInfo(c, user.ID)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			ErrorMessageWithErr("failed to fetch profile info", err),
		)
		return
	}
	c.JSON(http.StatusOK, userInfo)
}

// NewUserService creates a new instance of UserService satisfying UserIface
func NewUserService(cfg *config.AppConfig) UserIface {
	repo, err := database.NewUserRepository(cfg)
	if err != nil {
		utils.Logger.Fatal("failed to get user repository client", zap.Error(err))
		panic(err)
	}
	return &UserService{repo: repo}
}
