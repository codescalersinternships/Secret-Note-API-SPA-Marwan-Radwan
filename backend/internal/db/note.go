package db

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	ID             string    `gorm:"type:text;primaryKey" json:"id"`
	Content        string    `json:"content" gorm:"context"`
	RemainingViews int64     `json:"remaining_views" gorm:"remaining_views"`
	ExpireAt       time.Time `json:"expire_at" gorm:"expire_at"`
	CreatedAt      time.Time `json:"created_at" gorm:"expire_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"expire_at"`
}

func (note *Note) BeforeCreate(tx *gorm.DB) (err error) {
	note.ID = uuid.New().String()
	encryptedContent, err := encrypt(note.Content)
	if err != nil {
		return err
	}
	note.Content = encryptedContent
	return
}

func (note *Note) BeforeUpdate(tx *gorm.DB) (err error) {
	encryptedContent, err := encrypt(note.Content)
	if err != nil {
		return err
	}
	note.Content = encryptedContent
	return
}

func (note *Note) AfterFind(tx *gorm.DB) (err error) {
	decryptedContent, err := decrypt(note.Content)
	if err != nil {
		return err
	}
	note.Content = decryptedContent
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

func encrypt(text string) (string, error) {
	encryptionKey := []byte(os.Getenv("ENCRYPTION_KEY"))
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}
	plaintext := []byte(text)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(encryptedText string) (string, error) {
	encryptionKey := []byte(os.Getenv("ENCRYPTION_KEY"))
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
