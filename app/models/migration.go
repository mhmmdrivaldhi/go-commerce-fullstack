package models

import (
	"log"

	"gorm.io/gorm"
)

var ModelList = []interface{}{
	&User{},
	&Address{},
	&Product{},
	&Category{},
	&ProductImage{},
	&Section{},
	&Order{},
	&OrderItem{},
	&OrderCustomer{},
	&Payment{},
	&Shipment{},
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