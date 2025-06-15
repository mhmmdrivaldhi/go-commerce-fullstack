package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID        string  `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	Order     Order   `json:"order"`
	OrderID   string  `gorm:"size:36;index" json:"order_id"`
	Product   Product `json:"product"`
	ProductID string  `gorm:"size:36;index" json:"product_id"`
	Qty       int     `json:"qty"`
	BasePrice decimal.Decimal `gorm:"type:decimal(16,2)" json:"base_price"`
	BaseTotal decimal.Decimal `gorm:"type:decimal(16,2)" json:"base_total"`
	TaxAmount decimal.Decimal `gorm:"type:decimal(16,2)" json:"tax_amount"`
	TaxPercent decimal.Decimal `gorm:"type:decimal(10,2)" json:"tax_percent"`
	DiscountAmount decimal.Decimal `gorm:"type:decimal(16,2)" json:"discount_amount"`
	DiscountPercent decimal.Decimal `gorm:"type:decimal(10,2)" json:"discount_percent"`
	Subtotal decimal.Decimal `gorm:"type:decimal(16,2)" json:"subtotal"`
	Sku string `gorm:"size:36;index" json:"sku"`
	Name string `gorm:"size:255" json:"name"`
	Weight decimal.Decimal `gorm:"type:decimal(10,2)" json:"weight"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}