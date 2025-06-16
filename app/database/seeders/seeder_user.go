package seeders

import (
	"fmt"

	"github.com/mhmmdrivaldhi/golang_ecommerce/app/database/fakers"
	"gorm.io/gorm"
)

type UserSeeder struct {
	LastName string
	Seeder interface{}
}

func UserSeederRegister(db *gorm.DB) []UserSeeder{
	user := fakers.UserFaker(db)

	return []UserSeeder{
		{
			LastName: user.LastName,
			Seeder: user,
		},
	}
}

func UserSeed(db *gorm.DB) error {
	for _, users := range UserSeederRegister(db) {
		fmt.Printf("Running seeder Name %s\n", users.LastName)

		err := db.Debug().Create(users.Seeder).Error
		if err != nil {
			return fmt.Errorf("failed to seeding Name %s: %w", users.LastName, err)
		}
	}
	fmt.Println("User Seeding Completed Successfully!")
	return nil
}