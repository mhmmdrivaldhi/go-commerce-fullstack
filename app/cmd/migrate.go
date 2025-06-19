package cmd

import (
	"fmt"

	"github.com/mhmmdrivaldhi/golang_ecommerce/app/config"
	"github.com/mhmmdrivaldhi/golang_ecommerce/app/database/migration"
	"github.com/urfave/cli"
)

var MigrateCommand = cli.Command{
	Name: "db:migrate",
	Usage: "Migrate all database tables",
	Action: func (c *cli.Context) error {
		cfg := config.LoadConfig()
		db := config.InitDB(cfg.DB)

		migration.Migrate(db)
		fmt.Println("Migration Executed Successfully")
		return nil
	},
}