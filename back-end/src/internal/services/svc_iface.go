package services

import "github.com/gin-gonic/gin"

type AuthIface interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
	AllUserGuids(ctx *gin.Context)
	DeleteUser(ctx *gin.Context, userId string)
}

type NoteIface interface {
	CreateNote(*gin.Context)
	UpdateNote(*gin.Context)
	GetAllNotes(*gin.Context)
	GetNote(*gin.Context)
}
