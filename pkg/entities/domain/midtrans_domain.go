package domain

import "errors"

var (
	MessageSuccessWebhook           = "Webhook processed successfully"
	MessageSuccessCreateTransaction = "Transaction processed successfully"

	MessageFailedCreateTransaction = "failed to create transaction"

	ErrCreateTransactionFailed = errors.New("create transaction failed")
	ErrInvalidSignature        = errors.New("Invalid signature")
	ErrTransactionNotFound     = errors.New("Transaction not found")
)

type (
	MidtransPaymentRequest struct {
		Amount int64  `json:"amount" validate:"required"`
		Name   string `json:"name" validate:"required"`
		Email  string `json:"email" validate:"required"`
	}

	MidtransInvoiceUrl struct {
		Invoice string `json:"invoice"`
	}

	MidtransWebhookRequest struct {
		TransactionStatus string `json:"transaction_status"`
		OrderID           string `json:"order_id"`
		GrossAmount       string `json:"gross_amount"`
		FraudStatus       string `json:"fraud_status"`
		StatusCode        string `json:"status_code"`
		SignatureKey      string `json:"signature_key"`
	}

	MidtransWebhookResponse struct {
		TransactionStatus string `json:"transaction_status"`
		OrderID           string `json:"order_id"`
	}
)
