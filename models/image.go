package models

import "time"

// Image struct represents the image entity in the database
type Image struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	BookID     uint      `json:"book_id"`
	ImageURL   string    `json:"image_url"`
	UploadedAt time.Time `json:"uploaded_at"`
}
