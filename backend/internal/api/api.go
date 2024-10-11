package api

import (
	"bnot/backend/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(auth.AuthMiddleware())
	{
		api.GET("/notes", GetNotes)
		api.POST("/notes", CreateNote)
		api.GET("/notes/:id", GetNote)
		api.PUT("/notes/:id", UpdateNote)
		api.DELETE("/notes/:id", DeleteNote)
		api.GET("/notes/:id/versions", GetNoteVersions)
		api.POST("/notes/:id/versions", CreateNoteVersion)
	}
}
