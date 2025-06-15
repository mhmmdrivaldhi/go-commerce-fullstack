package models

import "time"

type Section struct {
	ID        string `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	Name      string `gorm:"size:100" json:"name"`
	Slug      string `gorm:"size:255" json:"slug"`
	Categories []Category `json:"categories"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}