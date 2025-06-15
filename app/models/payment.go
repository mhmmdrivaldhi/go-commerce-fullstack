package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Payment struct {
	ID      string `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	Order   Order  `json:"order"`
	OrderID string `gorm:"size:36;index" json:"order_id"`
	Number  string `gorm:"size:10;index" json:"number"`
	Amount  decimal.Decimal `gorm:"type:decimal(16,2)" json:"amount"`
	Method  string `gorm:"size:100" json:"method"`
	Status  string `gorm:"size:100" json:"status"`
	Token 	string `gorm:"size:100;index" json:"token"`
	Payload string `gorm:"type:text" json:"payload"`
	PaymentType string `gorm:"size:100" json:"payment_type"`
	VaNumber string `gorm:"size:100" json:"virtual_account_number"`
	BillCode string `gorm:"size:100" json:"bill_code"`
	BillKey string 	`gorm:"size:100" json:"bill_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}