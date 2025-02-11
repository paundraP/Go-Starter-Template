package subscription

import (
	"Go-Starter-Template/pkg/entities"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerKey string `mapstructure:"SERVER_KEY"`
	ClientKey string `mapstructure:"CLIENT_KEY"`
	IsProd    bool   `mapstructure:"IS_PROD"`
}

func NewMidtransConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func LogTransaction(transaction entities.Transaction) {
	logFile, err := os.OpenFile(
		"./logs/payments.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Println("Failed to open log file:", err)
		return
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Println("Failed to close log file:", err)
			return
		}
	}(logFile)

	logger := log.New(logFile, "", log.LstdFlags)

	// Log the payment info
	logger.Printf(
		"✅ [PAID] Invoice: %s | UserID: %d | Status: %s | Time: %s",
		transaction.Invoice,
		transaction.UserID,
		transaction.Status,
		time.Now().Format(time.RFC3339),
	)

	fmt.Println("Log entry successfully written")
}
