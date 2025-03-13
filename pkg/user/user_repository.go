package user

import (
	"Go-Starter-Template/entities"
	"context"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		RegisterUser(ctx context.Context, req entities.User) (entities.User, error)
		CheckUserByEmail(ctx context.Context, email string) bool
		GetUserByEmail(ctx context.Context, email string) (entities.User, error)
		CheckUserByID(ctx context.Context, id string) bool
		UpdateSubscriptionStatus(ctx context.Context, userID string) error
		UpdateProfile(ctx context.Context, user entities.User) error
		UpdateEducation(ctx context.Context, req entities.UserEducation) error
		PostExperience(ctx context.Context, req entities.UserExperience) error
		UpdateExperience(ctx context.Context, req entities.UserExperience) error
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

func (r *userRepository) CheckUserByEmail(ctx context.Context, email string) bool {
	var user entities.User
	if err := r.db.WithContext(ctx).First(user, "email = ?", email).Error; err != nil {
		return false
	}
	if user.Email != email {
		return false
	}
	return true
}
func (r *userRepository) CheckUserByID(ctx context.Context, id string) bool {
	var user entities.User
	if err := r.db.WithContext(ctx).First(user, "id = ?", id).Error; err != nil {
		return false
	}
	if user.ID.String() != id {
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

func (r *userRepository) UpdateProfile(ctx context.Context, user entities.User) error {
	if err := r.db.WithContext(ctx).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateEducation(ctx context.Context, req entities.UserEducation) error {
	if err := r.db.WithContext(ctx).Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) PostExperience(ctx context.Context, req entities.UserExperience) error {
	if err := r.db.WithContext(ctx).Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateExperience(ctx context.Context, req entities.UserExperience) error {
	if err := r.db.WithContext(ctx).Updates(&req).Error; err != nil {
		return err
	}
	return nil
}
