package cmd

import (
	"fmt"

	"github.com/mhmmdrivaldhi/golang_ecommerce/app/config"
	"github.com/mhmmdrivaldhi/golang_ecommerce/app/database/seeders"
	"github.com/urfave/cli"
)

var SeedCommand = cli.Command{
	Name: "db:seed",
	Usage: "Seed the database with initial data",
	Action: func (c *cli.Context) error  {
		cfg := config.LoadConfig()
		db := config.InitDB(cfg.DB)

		seeders.UserSeed(db)
		seeders.ProductSeed(db)
		fmt.Println("Seeding data has completed")
		return nil
	},
}