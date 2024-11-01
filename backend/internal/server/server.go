package server

import (
	"fmt"

	"github.com/codescalersinternships/Secret-Note-API-SPA-Marwan-Radwan.git/internal/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	DB     *db.DbClient
}

func NewServer(db *db.DbClient) *Server {
	r := gin.Default()
	s := Server{
		DB:     db,
		Router: r,
	}
	return &s
}

// Run launch server
func (s *Server) Run(port uint) error {
	s.registerRoutes()
	return s.Router.Run(fmt.Sprintf(":%d", 3000))
}

func (s *Server) registerRoutes() {
	r := s.Router

	routes := r.Group("/api")
	{
		note := routes.Group("/notes")
		{
			note.GET("/", s.GetNotes)
			note.POST("/", s.CreateNote)
			note.GET("/:id", s.GetNoteByID)
			note.PUT("/:id", s.UpdateNoteByID)
			note.DELETE("/:id", s.DeleteNoteByID)
		}
	}

}
