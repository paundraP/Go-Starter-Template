package rank

import (
	"Go-Starter-Template/domain"
	"Go-Starter-Template/entities"
	"Go-Starter-Template/internal/utils/pagination"
	"Go-Starter-Template/pkg/utility"
	"context"
	"gorm.io/gorm"
)

type (
	RankRepository interface {
		GetLeaderboard(ctx context.Context, metaReq pagination.Meta) (domain.GetLeaderboardRepository, error)
	}

	rankRepository struct {
		db *gorm.DB
	}
)

func NewRankRepository(db *gorm.DB) RankRepository {
	return &rankRepository{
		db: db,
	}
}

func (r *rankRepository) GetLeaderboard(ctx context.Context, metaReq pagination.Meta) (domain.GetLeaderboardRepository, error) {
	var users []entities.User

	db := r.db.WithContext(ctx).Model(&entities.User{})
	if err := utility.WithFilters(db, &metaReq, utility.AddModels(entities.User{}, "users")).
		Find(&users).Error; err != nil {
		return domain.GetLeaderboardRepository{}, err
	}
	return domain.GetLeaderboardRepository{
		Leaderboard: users,
		Meta:        metaReq,
	}, nil
}
