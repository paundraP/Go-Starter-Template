package payment

import (
	"Go-Starter-Template/internal/utils"
	"Go-Starter-Template/pkg/entities"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransConfig struct {
	ClientKey string
	ServerKey string
	IsProd    bool
}

func LoadMidtransConfig() MidtransConfig {
	isProd := os.Getenv("IS_PROD")
	prodMode := isProd == "true"

	return MidtransConfig{
		ClientKey: utils.GetEnv("CLIENT_KEY"),
		ServerKey: utils.GetEnv("SERVER_KEY"),
		IsProd:    prodMode,
	}
}

func LogTransaction(transaction entities.Transaction) {
	logFile, err := os.OpenFile(
		"./logs/payments.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
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

func NewMidtransClient() snap.Client {
	midtransConfig := LoadMidtransConfig()

	var client snap.Client
	client.New(midtransConfig.ServerKey, midtrans.Sandbox)
	if midtransConfig.IsProd {
		client.New(midtransConfig.ServerKey, midtrans.Production)
	}
	return client
}
