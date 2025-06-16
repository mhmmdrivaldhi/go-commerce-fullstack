package seeders

import (
	"fmt"

	"github.com/mhmmdrivaldhi/golang_ecommerce/app/database/fakers"
	"gorm.io/gorm"
)

type ProductSeeder struct {
	Name   string
	Seeder interface{}
}

func ProductSeederRegister(db *gorm.DB) []ProductSeeder {
	product := fakers.ProductFaker(db)
	
	return []ProductSeeder{
		{
			Name: product.Name,
			Seeder: product,
		},
	}
}

func ProductSeed(db *gorm.DB) error {
	for _, products := range ProductSeederRegister(db) {
		fmt.Printf("Running Product Seeder Name %s\n", products.Name)

		err := db.Debug().Create(products.Seeder).Error
		if err != nil {
			return fmt.Errorf("failed to product seeding Name %s: %w", products.Name, err)
		}
	}
	fmt.Println("Product Seeding Completed Successfully!")
	return nil
}