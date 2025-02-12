package payment

import (
	paymentconf "Go-Starter-Template/internal/config/paymentConf"
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func NewMidtransClient() snap.Client {
	midtransConfig, err := paymentconf.NewMidtransConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(midtransConfig.ServerKey)

	var client snap.Client
	client.New(midtransConfig.ServerKey, midtrans.Sandbox)
	if midtransConfig.IsProd {
		client.New(midtransConfig.ServerKey, midtrans.Production)
	}
	return client
}
