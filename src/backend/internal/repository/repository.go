package repository

import (
	"context"

	"bot-profile-manager/internal/models"
)

type ProfileRepository interface {
	GetProfile(ctx context.Context, botID string) (*models.Profile, error)
	GetAllProfiles(ctx context.Context) ([]*models.Profile, error)
	UpsertProfile(ctx context.Context, profile *models.Profile) error
}
