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

func (v *VanillaServer) Initialize() {
	// service - entry to business logic
	utils.InitLogging()
	utils.Logger.Info("initializing server...")
	noteSvc := NewGNoteService(v.cfg)
	authSvc := NewAuthService(v.cfg)

	v.router.Use(CORSMiddleware())

	public := v.router.Group("/auth")
	public.POST("/login", authSvc.Login)
	public.POST("/register", authSvc.Register)

	private := v.router.Group("/user")
	private.Use(AuthMiddleware())
	private.POST("/note", noteSvc.CreateNote)
	private.GET("/notes", noteSvc.GetAllNotes)
	private.GET("/note/:noteId", noteSvc.GetNote)
	private.PUT("/note/:noteId", noteSvc.UpdateNote)
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
