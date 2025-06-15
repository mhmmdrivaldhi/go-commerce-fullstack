package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string 		 `gorm:"size:36;not null;uniqueIndex;primary_key" json:"id"`
	FirstName     string 		 `gorm:"size:100;not null" json:"first_name"`
	LastName      string 		 `gorm:"size:100;not null" json:"last_name"`
	Adresses	  []Address 	 `json:"address"`
	Email         string 		 `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Password      string 		 `gorm:"size:255;not null;" json:"password"`
	RememberToken string 		 `gorm:"size:255;not null" json:"remember_token"`
	CreatedAt     time.Time 	 `json:"created_at"`
	UpdatedAt     time.Time 	 `json:"updated_at"`
	DeletedAt	  gorm.DeletedAt `json:"deleted_at"`
}