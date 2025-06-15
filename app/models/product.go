package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID       string `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	ParentID string `gorm:"size:36;index" json:"parent_id"`
	User     User   `json:"user"`
	UserID   string `gorm:"size:36;index" json:"user_id"`
	ProductImages []ProductImage `json:"product_images"`
	Categories []Category `gorm:"many2many:product_categories" json:"category"`
	Sku      string `gorm:"size:100;index" json:"sku"`
	Name     string `gorm:"size:255;index" json:"name"`
	Slug     string `gorm:"size:255;index" json:"slug"`
	Price    decimal.Decimal `gorm:"type:decimal(16,2);" json:"price"`
	Stock	 int	`json:"stock"`
	Weight	 decimal.Decimal `gorm:"type:decimal(10,2)" json:"weight"`
	ShortDescription string `gorm:"size:255" json:"short_description"`
	Description string `gorm:"type:text" json:"description"`
	Status int `gorm:"default:0" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}