package services

import "src/pkg/utils"

func (v *VanillaServer) Initialize() {
	// service - entry to business logic
	utils.InitLogging()
	utils.Logger.Info("initializing server...")
	noteSvc := NewGNoteService(v.cfg)
	authSvc := NewAuthService(v.cfg)
	userSvc := NewUserService(v.cfg)

	v.router.Use(CORSMiddleware())

	public := v.router.Group("/auth")
	public.POST("/login", authSvc.Login)
	public.POST("/register", authSvc.Register)

	private := v.router.Group("/user")
	private.Use(AuthMiddleware())

	private.GET("/info", userSvc.Info)
	private.POST("/note", noteSvc.CreateNote)
	private.GET("/notes", noteSvc.GetAllNotes)
	private.GET("/note/:noteId", noteSvc.GetNote)
	private.PUT("/note/:noteId", noteSvc.UpdateNote)
}
