package config

import (
	"Go-Starter-Template/internal/api/handlers"
	"Go-Starter-Template/internal/api/routes"
	"Go-Starter-Template/internal/middleware"
	"Go-Starter-Template/internal/utils"
	"Go-Starter-Template/internal/utils/storage"
	"Go-Starter-Template/pkg/jwt"
	"Go-Starter-Template/pkg/midtrans"
	"Go-Starter-Template/pkg/rank"
	"Go-Starter-Template/pkg/user"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
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
	err = os.MkdirAll("./logs", os.ModePerm)
	if err != nil {
		log.Fatalf("error creating logs directory: %v", err)
	}
	file, err := os.OpenFile(
		"./logs/app.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Output:     file,
	}))

	// Uncomment if you want to use limiter to your endpoint
	//app.Use(limiter.New(limiter.Config{
	//	Max:        10,
	//	Expiration: 1 * time.Second,
	//}))

	// utils
	s3 := storage.NewAwsS3()

	// Repository
	userRepository := user.NewUserRepository(db)
	midtransRepository := midtrans.NewMidtransRepository(db)
	rankRepository := rank.NewRankRepository(db)

	// Service
	jwtService := jwt.NewJWTService()
	userService := user.NewUserService(userRepository, jwtService, s3)
	rankService := rank.NewRankService(rankRepository, userRepository)
	midtransService := midtrans.NewMidtransService(
		midtransRepository,
		userRepository,
	)

	// Handler
	userHandler := handlers.NewUserHandler(userService, validator)
	midtransHandler := handlers.NewMidtransHandler(midtransService, validator)
	rankHandler := handlers.NewRankHandler(rankService, validator)

	// routes
	routesConfig := routes.Config{
		App:             app,
		UserHandler:     userHandler,
		MidtransHandler: midtransHandler,
		RankHandler:     rankHandler,
		Middleware:      middlewares,
		JWTService:      jwtService,
	}
	routesConfig.Setup()
	return app, nil
}
