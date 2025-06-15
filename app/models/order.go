package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	ID         string      `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	UserID     string      `gorm:"size:36;index"`
	User       User        `json:"user"`
	OrderItems []OrderItem `json:"order_items"`
	OrderCustomer *OrderCustomer `json:"order_customer"`
	Code      string `gorm:"size:50;index" json:"code"`
	Status    int    `json:"status"`
	OrderDate time.Time `json:"order_date"`
	PaymentDue time.Time `json:"payment_due"`
	PaymentStatus string `gorm:"size:50;index" json:"payment_status"`
	PaymentToken string `gorm:"size:100" json:"payment_token"`
	TotalPrice decimal.Decimal `gorm:"type:decimal(16,2)" json:"total_price"`
	TaxAmount decimal.Decimal `gorm:"type:decimal(16,2)" json:"tax_amount"`
	TaxPercent decimal.Decimal `gorm:"type:decimal(10,2)" json:"tax_percent"`
	DiscountAmount decimal.Decimal `gorm:"type:decimal(16,2)" json:"discount_amount"`
	DiscountPercent decimal.Decimal `gorm:"type:decimal(10,2)" json:"discount_percent"`
	ShippingCost decimal.Decimal `gorm:"type:decimal(16,2)" json:"shipping_cost"`
	GrandTotal decimal.Decimal `gorm:"type:decimal(16,2)" json:"grand_total"`
	Note string `gorm:"type:text" json:"note"`
	ShippingCourier string `gorm:"size:100" json:"shipping_courier"`
	ShippingServiceName string `gorm:"size:100" json:"shipping_service_name"`
	ApprovedBy string `gorm:"size:100" json:"approved_by"`
	ApprovedAt time.Time `json:"approved_at"`
	CancelledBy string `gorm:"size:100" json:"cancelled_by"`
	CancelledAt time.Time `json:"cancelled_at"`
	CancellationNote string `gorm:"size:255" json:"cancellation_note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}