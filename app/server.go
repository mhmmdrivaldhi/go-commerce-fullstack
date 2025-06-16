package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mhmmdrivaldhi/golang_ecommerce/app/config"
	"github.com/mhmmdrivaldhi/golang_ecommerce/app/database/migration"
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

	// DB Configuration
	cfg := config.LoadConfig()
	db := config.InitDB(cfg.DB)

	// Models Migration
	migration.Migrate(db) 

	var server = Server{}
	server.Initialize(db, *cfg.App)
	server.Run(":" + cfg.App.AppPort)
}