package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mhmmdrivaldhi/golang_ecommerce/app/cmd"
	"github.com/mhmmdrivaldhi/golang_ecommerce/app/config"
	// "github.com/mhmmdrivaldhi/golang_ecommerce/app/database/seeders"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(db *gorm.DB, appConfig config.AppConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.DB = db
	server.Routes()

	// seeders.UserSeed(server.DB)
	// seeders.ProductSeed(server.DB)
}

func (server *Server) Run(addr string) {
	log.Printf("Server is running on port%s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load .env file: %v", err)
	}

	cmdApp := &cli.App{
		Name: "go-commerce",
		Usage: "Command Line Interface for Go-Commerce",
		Commands: []cli.Command{
			cmd.MigrateCommand,
			cmd.SeedCommand,
			{
				Name: "serve",
				Usage: "Run the HTTP server",
				Action: func (c *cli.Context) error {
					// Load Config and DB
					cfg := config.LoadConfig()
					db := config.InitDB(cfg.DB)

					var server = Server{}
					
					server.Initialize(db, *cfg.App)
					server.Run(":" + cfg.App.AppPort)
					return nil
				},
			},
		},
	}
	if err := cmdApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}