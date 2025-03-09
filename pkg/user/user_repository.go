package user

import (
	"Go-Starter-Template/pkg/entities"
	"context"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		RegisterUser(ctx context.Context, req entities.User) (entities.User, error)
		CheckUser(ctx context.Context, email string) bool
		GetUserByEmail(ctx context.Context, email string) (entities.User, error)
		UpdateSubscriptionStatus(ctx context.Context, userID string) error
	}
	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) RegisterUser(ctx context.Context, req entities.User) (entities.User, error) {
	if err := r.db.WithContext(ctx).Create(&req).Error; err != nil {
		return entities.User{}, err
	}
	return req, nil
}

func (r *userRepository) CheckUser(ctx context.Context, email string) bool {
	var user entities.User
	if err := r.db.WithContext(ctx).First(user, "email = ?", email).Error; err != nil {
		return false
	}
	if user.Email != email {
		return false
	}
	return true
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateSubscriptionStatus(ctx context.Context, userID string) error {
	if err := r.db.WithContext(ctx).
		Model(&entities.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{"subscribe": true}).Error; err != nil {
		return err
	}
	return nil
}
