package user

import (
	entities2 "Go-Starter-Template/entities"
	"context"
	"errors"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CreateUser(ctx context.Context, user *entities2.User) error
		GetEmail(ctx context.Context, email string) (*entities2.User, error)
		UpdateUser(ctx context.Context, user entities2.User) (*entities2.User, error)
		GetUserByID(ctx context.Context, id string) (*entities2.User, error)
		GetRankByTotalPoint(ctx context.Context, totalPoint int) (*entities2.Rank, error)
		UpdateSubscriptionStatus(ctx context.Context, userID string) error
	}
	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entities2.User) error {

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetEmail(ctx context.Context, email string) (*entities2.User, error) {
	var user entities2.User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user entities2.User) (*entities2.User, error) {
	if err := r.db.WithContext(ctx).Updates(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetUserByID(ctx context.Context, id string) (*entities2.User, error) {
	var user entities2.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetRankByTotalPoint(ctx context.Context, totalPoint int) (*entities2.Rank, error) {
	var rank entities2.Rank
	if err := r.db.WithContext(ctx).First(&rank, "lower_point <= ? AND upper_point >= ?", totalPoint, totalPoint).Error; err != nil {
		return nil, err
	}
	return &rank, nil
}

func (r *userRepository) UpdateSubscriptionStatus(ctx context.Context, userID string) error {
	if err := r.db.WithContext(ctx).
		Model(&entities2.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{"subscribe": true}).Error; err != nil {
		return err
	}
	return nil
}
