package midtrans

import (
	"Go-Starter-Template/entities"
	"context"
	"gorm.io/gorm"
)

type (
	MidtransRepository interface {
		CreateTransaction(transaction entities.Transaction) error
		GetOrderID(ctx context.Context, orderID string) (entities.Transaction, error)
		UpdateTransaction(ctx context.Context, transaction entities.Transaction) error
	}

	midtransRepository struct {
		db *gorm.DB
	}
)

func NewMidtransRepository(db *gorm.DB) MidtransRepository {
	return &midtransRepository{db}
}

func (r *midtransRepository) CreateTransaction(transaction entities.Transaction) error {
	return r.db.Create(&transaction).Error
}

func (r *midtransRepository) GetOrderID(ctx context.Context, orderID string) (entities.Transaction, error) {
	var transaction entities.Transaction
	if err := r.db.WithContext(ctx).First(&transaction, "order_id = ?", orderID).Error; err != nil {
		return entities.Transaction{}, err
	}
	return transaction, nil
}

func (r *midtransRepository) UpdateTransaction(ctx context.Context, transaction entities.Transaction) error {
	return r.db.WithContext(ctx).Save(&transaction).Error
}
