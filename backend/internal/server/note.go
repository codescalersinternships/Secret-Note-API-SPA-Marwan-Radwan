package server

import (
	"log"
	"net/http"

	"github.com/codescalersinternships/Secret-Note-API-SPA-Marwan-Radwan.git/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Server) GetNotes(ctx *gin.Context) {
	dbClient := s.DB
	notes, err := db.GetNotes(dbClient.DB)
	if err != nil {
		log.Printf("Error finding all notes , %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notes)
}

func (s *Server) GetNoteByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, _ := uuid.Parse(id)
	dbClient := s.DB
	note, err := db.GetNoteById(dbClient.DB, uuid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	if note.RemainingViews == 0 {
		if err = db.DeleteNote(dbClient.DB, uuid); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
	}
	note.RemainingViews = note.RemainingViews - 1
	if err = db.UpdateNote(dbClient.DB, note); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	ctx.JSON(http.StatusOK, note)
}

func (s *Server) CreateNote(ctx *gin.Context) {
	var note db.Note
	dbClient := s.DB
	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.CreateNote(dbClient.DB, &note); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, note)
}

func (s *Server) UpdateNoteByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var note *db.Note
	uuid, _ := uuid.Parse(id)
	dbClient := s.DB
	note, err := db.GetNoteById(dbClient.DB, uuid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	if err := ctx.ShouldBindJSON(note); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.UpdateNote(dbClient.DB, note); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, note)
}

func (s *Server) DeleteNoteByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, _ := uuid.Parse(id)
	dbClient := s.DB

	if err := db.DeleteNote(dbClient.DB, uuid); err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
