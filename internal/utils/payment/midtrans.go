package payment

import (
	paymentconf "Go-Starter-Template/internal/config/payment_config"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func NewMidtransClient() snap.Client {
	midtransConfig := paymentconf.LoadMidtransConfig()

	var client snap.Client
	client.New(midtransConfig.ServerKey, midtrans.Sandbox)
	if midtransConfig.IsProd {
		client.New(midtransConfig.ServerKey, midtrans.Production)
	}
	return client
}
