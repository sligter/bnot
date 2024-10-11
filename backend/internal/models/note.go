package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserID   uint   `json:"user_id"`
	Versions []Version
}

type Version struct {
	gorm.Model
	NoteID  uint   `json:"note_id"`
	Content string `json:"content"`
}
