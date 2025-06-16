package fakers

import (
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/mhmmdrivaldhi/golang_ecommerce/app/models"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID: uuid.New().String(),
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Email: faker.Email(),
		Password: "$2a$12$P5FWGa08C.amVIztHjCdG.JLh1chCVj4ljvlWGLUgVSEPpXeeqodS",
		RememberToken: "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}