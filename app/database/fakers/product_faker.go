package fakers

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/mhmmdrivaldhi/golang_ecommerce/app/models"
	"github.com/mhmmdrivaldhi/golang_ecommerce/app/utils"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ProductFaker(db *gorm.DB) *models.Product {
	user := UserFaker(db)
	err := db.Create(user).Error
	if err != nil {
		log.Fatal(err)
	}

	name := faker.Name()
	slug_title := slug.Make(name)

	return &models.Product{
		ID: uuid.New().String(),
		UserID: user.ID,
		Sku: utils.GenerateBarcodeSKU(),
		Name: name,
		Slug: slug_title,
		Price: decimal.NewFromFloat(fakePrice()),
		Stock: rand.Intn(100),
		Weight: decimal.NewFromFloat(rand.Float64()),
		ShortDescription: faker.Paragraph(),
		Description: faker.Paragraph(),
		Status: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}
}

func precision(value float64, pre int) float64 {
	div := math.Pow10(pre)
	return float64(int64(value * div))
}

func fakePrice() float64 {
	return precision(rand.Float64() * math.Pow10(rand.Intn(8)), rand.Intn(2) + 1)
}