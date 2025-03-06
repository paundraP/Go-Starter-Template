package user

import (
	"Go-Starter-Template/pkg/entities"
	"context"
	"errors"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CreateUser(ctx context.Context, user *entities.User) error
		GetEmail(ctx context.Context, email string) (*entities.User, error)
		UpdateUser(ctx context.Context, user entities.User) (*entities.User, error)
		GetUserByID(ctx context.Context, id string) (*entities.User, error)
		UpdateSubscriptionStatus(ctx context.Context, userID string) error
	}
	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entities.User) error {

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user entities.User) (*entities.User, error) {
	if err := r.db.WithContext(ctx).Updates(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetUserByID(ctx context.Context, id string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
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
