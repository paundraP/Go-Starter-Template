package config

import (
	"Go-Starter-Template/internal/api/handlers"
	"Go-Starter-Template/internal/api/routes"
	"Go-Starter-Template/internal/middleware"
	"Go-Starter-Template/internal/utils"
	"Go-Starter-Template/pkg/midtrans"
	"Go-Starter-Template/pkg/user"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

func NewApp(db *gorm.DB) (*fiber.App, error) {
	utils.InitValidator()
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
	middlewares := middleware.NewMiddleware()
	validator := utils.Validate

	// load all env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// setting up logging and limiter
	logDir := "logs"
	logFile := "app.log"
	log_path := filepath.Join(logDir, logFile)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(log_path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("error closing file: %v", err)
		}
	}(file)

	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Output:     file,
	}))
	// Repository
	userRepository := user.NewUserRepository(db)
	midtransRepository := midtrans.NewMidtransRepository(db)

	// Service
	userService := user.NewUserService(userRepository)
	midtransService := midtrans.NewMidtransService(
		midtransRepository,
		userRepository,
	)

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
