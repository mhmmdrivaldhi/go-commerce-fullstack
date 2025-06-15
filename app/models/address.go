package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID            string `gorm:"size:100;not null;uniqueIndex;primary_key" json:"id"`
	User          User	 `json:"user"`
	UserID		  string `gorm:"size:36;index"`
	FirstName     string `gorm:"size:100" json:"first_name"`
	LastName      string `gorm:"size:100" json:"last_name"`
	IsPrimary     bool   `json:"is_primary"`
	FirstAddress  string `gorm:"size:100" json:"first_address"`
	SecondAddress string `gorm:"size:100" json:"second_address"`
	CityID		  string `gorm:"size:10" json:"city_id"`
	ProvinceID	  string `gorm:"size:10" json:"province_id"`
	PostCode      string `gorm:"size:10" json:"post_code"`
	Phone         string `gorm:"size:50" json:"phone_number"`
	Email         string `gorm:"size:100" json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt	  gorm.DeletedAt `json:"deleted_at"`
}