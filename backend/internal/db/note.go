package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	ID             string    `gorm:"type:text;primaryKey" json:"id"`
	Content        string    `json:"content"` // Store encrypted data directly in this field
	RemainingViews int64     `json:"remaining_views"`
	ExpireAt       time.Time `json:"expire_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Encrypt content before saving to the database
func (note *Note) BeforeCreate(tx *gorm.DB) (err error) {
	note.ID = uuid.New().String()
	return
}

func CreateNote(db *gorm.DB, note *Note) error {
	return db.Create(note).Error
}

func GetNoteById(db *gorm.DB, id uuid.UUID) (*Note, error) {
	var note Note
	if err := db.First(&note, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &note, nil
}

func GetNotes(db *gorm.DB) ([]Note, error) {
	var notes []Note
	if err := db.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func UpdateNote(db *gorm.DB, note *Note) error {
	return db.Save(note).Error
}

func DeleteNote(db *gorm.DB, id uuid.UUID) error {
	return db.Delete(&Note{}, "id = ?", id).Error
}
