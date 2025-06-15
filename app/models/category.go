package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        string    `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	ParentID  string    `gorm:"size:36" json:"parent_id"`
	Section   Section   `json:"section"`
	SectionID string    `gorm:"size:36;index" json:"section_id"`
	Products  []Product `gorm:"many2many:product_categories" json:"product_categories"`
	Name      string    `gorm:"size:100" json:"name"`
	Slug      string    `gorm:"size:100" json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}