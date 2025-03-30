package rank

import (
	"Go-Starter-Template/domain"
	"Go-Starter-Template/internal/utils/pagination"
	"Go-Starter-Template/pkg/user"
	"context"
)

type (
	RankService interface {
		Leaderboard(ctx context.Context, metaReq pagination.Meta) (domain.GetLeaderboardResponse, error)
	}

	rankService struct {
		rankRepository RankRepository
		userRepository user.UserRepository
	}
)

func NewRankService(rankRepository RankRepository, userRepository user.UserRepository) RankService {
	return &rankService{
		rankRepository: rankRepository,
		userRepository: userRepository,
	}
}

func (s *rankService) Leaderboard(ctx context.Context, metaReq pagination.Meta) (domain.GetLeaderboardResponse, error) {
	users, err := s.rankRepository.GetLeaderboard(ctx, metaReq)
	if err != nil {
		return domain.GetLeaderboardResponse{}, domain.ErrGetLeaderboard
	}

	var userResp []domain.UserRankDetail
	for _, user := range users.Leaderboard {
		totalPoint := user.ActivePoint + user.LevelPoint
		rank, err := s.userRepository.GetRankByTotalPoint(ctx, totalPoint)
		if err != nil {
			return domain.GetLeaderboardResponse{}, err
		}
		userResp = append(userResp, domain.UserRankDetail{
			ID:            user.ID,
			Name:          user.Name,
			ActivityPoint: user.ActivePoint,
			LevelPoint:    user.LevelPoint,
			TotalPoint:    totalPoint,
			Rank:          rank.Name,
		})
	}
	return domain.GetLeaderboardResponse{
		UserRank: userResp,
		Meta:     metaReq,
	}, nil
}
