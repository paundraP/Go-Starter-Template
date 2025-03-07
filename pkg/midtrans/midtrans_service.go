package midtrans

import (
	"Go-Starter-Template/internal/utils/payment"
	paymentconf "Go-Starter-Template/internal/utils/payment"
	"Go-Starter-Template/pkg/entities"
	"Go-Starter-Template/pkg/entities/domain"
	"Go-Starter-Template/pkg/user"
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"math/big"
	"os"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type (
	MidtransService interface {
		CreateTransaction(req domain.MidtransPaymentRequest, userID string) (domain.MidtransInvoiceUrl, error)
		MidtransWebHook(ctx context.Context, req domain.MidtransWebhookRequest) (domain.MidtransWebhookResponse, error)
	}

	midtransService struct {
		midtransRepository MidtransRepository
		userRepository     user.UserRepository
	}
)

func NewMidtransService(midtransRepo MidtransRepository, userRepository user.UserRepository) MidtransService {
	return &midtransService{
		midtransRepository: midtransRepo,
		userRepository:     userRepository,
	}
}

func validateSignature(orderID, statusCode, grossAmount, receivedSignature string) bool {
	serverKey := os.Getenv("SERVER_KEY")
	rawString := orderID + statusCode + grossAmount + serverKey

	hash := sha512.Sum512([]byte(rawString))
	expectedSignature := hex.EncodeToString(hash[:])

	return expectedSignature == receivedSignature
}

func (s *midtransService) CreateTransaction(req domain.MidtransPaymentRequest, userID string) (domain.MidtransInvoiceUrl, error) {
	client := payment.NewMidtransClient()
	orderID := GenerateRandomString()
	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: req.Amount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: req.Name,
			Email: req.Email,
		},
	}

	snapResp, er := client.CreateTransaction(request)
	if er != nil {
		return domain.MidtransInvoiceUrl{}, domain.ErrCreateTransactionFailed
	}

	userid, err := uuid.Parse(userID)
	if err != nil {
		return domain.MidtransInvoiceUrl{}, domain.ErrParseUUID
	}

	transact := entities.Transaction{
		ID:      uuid.New(),
		UserID:  userid,
		Status:  "pending",
		Invoice: snapResp.RedirectURL,
		OrderID: orderID,
	}

	if err = s.midtransRepository.CreateTransaction(transact); err != nil {
		return domain.MidtransInvoiceUrl{}, err
	}

	return domain.MidtransInvoiceUrl{
		Invoice: snapResp.RedirectURL,
	}, nil
}

func (s *midtransService) MidtransWebHook(ctx context.Context, req domain.MidtransWebhookRequest) (domain.MidtransWebhookResponse, error) {
	if !validateSignature(
		req.OrderID,
		req.StatusCode,
		req.GrossAmount,
		req.SignatureKey,
	) {
		return domain.MidtransWebhookResponse{}, domain.ErrInvalidSignature
	}

	transaction, err := s.midtransRepository.GetOrderID(ctx, req.OrderID)
	if err != nil {
		return domain.MidtransWebhookResponse{}, domain.ErrTransactionNotFound
	}

	transactionStatus := req.TransactionStatus
	fraudStatus := req.FraudStatus

	switch transactionStatus {
	case "capture":
		if fraudStatus == "accept" {
			transaction.Status = "paid"
			paymentconf.LogTransaction(transaction)
		} else {
			transaction.Status = "fraud"
		}
	case "settlement":
		transaction.Status = "paid"
		paymentconf.LogTransaction(transaction)
	case "deny", "cancel", "expire":
		transaction.Status = "failed"
	case "pending":
		transaction.Status = "pending"
	case "refund":
		transaction.Status = "refunded"
	}

	if err := s.midtransRepository.UpdateTransaction(ctx, transaction); err != nil {
		return domain.MidtransWebhookResponse{}, err
	}

	if err := s.userRepository.UpdateSubscriptionStatus(ctx, transaction.UserID.String()); err != nil {
		return domain.MidtransWebhookResponse{}, err
	}

	return domain.MidtransWebhookResponse{
		TransactionStatus: transaction.Status,
		OrderID:           transaction.Invoice,
	}, nil
}

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers = "0123456789"
)

func GenerateRandomString() string {
	result := make([]byte, 8)

	for i := 0; i < 4; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[num.Int64()]
	}

	for i := 4; i < 8; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(numbers))))
		result[i] = numbers[num.Int64()]
	}

	return string(result)
}
