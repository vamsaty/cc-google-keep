package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"src/internal/config"
	"src/pkg/utils"
)

type ServerIface interface {
	Run()
	Initialize()
}

type VanillaServer struct {
	cfg    *config.AppConfig
	router *gin.Engine
}

func (v *VanillaServer) Run() {
	utils.Logger.Info("starting server...")
	log.Fatalln(v.router.Run(v.cfg.ServerConfig.Address()))
}

func NewVanillaServerWithConfig(cfg *config.AppConfig) ServerIface {
	return &VanillaServer{
		cfg:    cfg,
		router: gin.New(),
	}
}

func NewVanillaServer() ServerIface {
	cfg := config.LoadAppConfig()
	return &VanillaServer{
		cfg:    &cfg,
		router: gin.New(),
	}
}
