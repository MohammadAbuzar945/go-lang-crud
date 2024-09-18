package models

import "time"

type Book struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"publishedAt"`
}
