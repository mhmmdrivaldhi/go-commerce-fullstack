package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	ID         string  `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	Product    Product `json:"product"`
	ProductID  string  `gorm:"size:36;index"`
	Path       string  `gorm:"type:text" json:"path"`
	ExtraLarge string  `gorm:"size:255" json:"extra_large"`
	Large      string  `gorm:"size:255" json:"large"`
	Medium     string  `gorm:"size:255" json:"medium"`
	Small      string  `gorm:"size:255" json:"small"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}