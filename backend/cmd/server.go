package main

import (
	"echocrud/internal/db"
	"echocrud/internal/handler"
	"echocrud/internal/repository"
	"echocrud/internal/router"
	"echocrud/internal/seeder"
	"echocrud/internal/service"
	"flag"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	runSeed := flag.Bool("seed", false, "Run to populate the database with initial data")
	flag.Parse()

	server := echo.New()

	//Middleware + CORS
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	// Database
	database := db.NewPostgresConnection()

	//seed
	if *runSeed {
		log.Println("Starting seeder execution...")
		seeder.Seed(database)
		log.Println("Seeder executed successfully. The application will now exit.")
		os.Exit(0) 
	}

	// Repositories
	EstablishmentRepository := repository.NewEstablishmentRepository(database)
	StoreRepository := repository.NewStoreRepository(database)

	// Services
	EstablishmentService := service.NewEstablishmentService(EstablishmentRepository)
	StoreService := service.NewStoreService(StoreRepository, EstablishmentRepository)

	// Handlers
	EstablishmentHandler := handler.NewEstablishmentHandler(EstablishmentService)
	StoreHandler := handler.NewStoreHandler(StoreService)

	//Routes
	router.RegisterRoutes(server, EstablishmentHandler, StoreHandler)


	log.Println("🚀 Server Running in /api/v1")
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" 
	}

	server.Logger.Fatal(server.Start(":" + port))
}
