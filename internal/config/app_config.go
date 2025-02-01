package config

import (
	"Go-Starter-Template/database"
	"Go-Starter-Template/database/seeder"
	"Go-Starter-Template/internal/api/handlers"
	"Go-Starter-Template/internal/api/routes"
	"Go-Starter-Template/internal/middleware"
	"Go-Starter-Template/internal/utils"
	"Go-Starter-Template/pkg/midtrans"
	"Go-Starter-Template/pkg/user"
	"errors"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func NewApp() (*fiber.App, error) {
	utils.InitValidator()
	app := fiber.New()

	middlewares := middleware.NewMiddleware()
	validator := utils.Validate

	// load all env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// setting up logging and limiter
	file, err := os.OpenFile("./logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Output:     file,
	}))

	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Second,
	}))

	// setting up database (migration and data)
	db := database.NewDB()
	if db == nil {
		return nil, errors.New("db is nil")
	}
	migrate := flag.Bool("migrate", false, "migrate db")
	seed := flag.Bool("seed", false, "seed db")
	flag.Parse()
	if *migrate {
		if err := database.Migrate(db); err != nil {
			return nil, err
		}
	}
	if *seed {
		if err := seeder.Seed(db); err != nil {
			return nil, err
		}
	}

	// Repository
	userRepository := user.NewUserRepository(db)
	midtransRepository := midtrans.NewMidtransRepository(db)

	// Service
	userService := user.NewUserService(userRepository)
	midtransService := midtrans.NewMidtransService(midtransRepository, userRepository)

	// Handler
	userHandler := handlers.NewUserHandler(userService, validator)
	midtransHandler := handlers.NewMidtransHandler(midtransService, validator)

	// routes
	routesConfig := routes.Config{
		App:             app,
		UserHandler:     userHandler,
		MidtransHandler: midtransHandler,
		Middleware:      middlewares,
	}
	routesConfig.Setup()
	return app, nil
}
