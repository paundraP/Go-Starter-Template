package domain

import (
	"Go-Starter-Template/entities"
	"Go-Starter-Template/internal/utils/pagination"
	"errors"
	"github.com/google/uuid"
)

var (
	MessageFailedGetLeaderboard = "failed to get the leaderboard"

	MessageSuccessGetLeaderboard = "success get the leaderboard"

	ErrGetLeaderboard = errors.New("error getting leaderboard")
)

type (
	UserRankDetail struct {
		ID            uuid.UUID `json:"id"`
		Name          string    `json:"name"`
		Username      string    `json:"username"`
		ActivityPoint int       `json:"activity_point"`
		LevelPoint    int       `json:"level_point"`
		TotalPoint    int       `json:"total_point"`
		Rank          string    `json:"rank"`
	}

	GetLeaderboardResponse struct {
		UserRank []UserRankDetail `json:"user_rank"`
		Meta     pagination.Meta  `json:"meta"`
	}
	GetLeaderboardRepository struct {
		Leaderboard []entities.User
		pagination.Meta
	}
)
