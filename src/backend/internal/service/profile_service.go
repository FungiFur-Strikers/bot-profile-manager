package service

import (
	"context"
	"time"

	"bot-profile-manager/internal/models"
	"bot-profile-manager/internal/repository"
)

type ProfileService struct {
	repo repository.ProfileRepository
}

func NewProfileService(repo repository.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetProfile(ctx context.Context, botID string) (interface{}, error) { // interfaceに変更
	if botID == "*" {
		return s.repo.GetAllProfiles(ctx)
	}
	return s.repo.GetProfile(ctx, botID)
}

func (s *ProfileService) UpsertProfile(ctx context.Context, profile *models.Profile) error {
	now := time.Now()
	if profile.CreatedAt.IsZero() {
		profile.CreatedAt = now
	}
	profile.UpdatedAt = now
	return s.repo.UpsertProfile(ctx, profile)
}
