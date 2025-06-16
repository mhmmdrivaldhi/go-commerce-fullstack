package migration

import (
	"log"

	"github.com/mhmmdrivaldhi/golang_ecommerce/app/models"
	"gorm.io/gorm"
)

var ModelList = []interface{}{

	&models.User{},
	&models.Address{},
	&models.Product{},
	&models.Category{},
	&models.ProductImage{},
	&models.Section{},
	&models.Order{},
	&models.OrderItem{},
	&models.OrderCustomer{},
	&models.Payment{},
	&models.Shipment{},
}

func Migrate(db *gorm.DB) {
	for _, model := range ModelList {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Migration failed for %T: %v", model, err)
		}
		log.Printf("Migrated: %T", model)
	}
	log.Println("Successfully migrated All models.")
}