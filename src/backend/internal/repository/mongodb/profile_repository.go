package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"bot-profile-manager/internal/models"
	"bot-profile-manager/internal/repository"
)

type profileRepository struct {
	collection *mongo.Collection
}

func NewProfileRepository(client *mongo.Client, dbName string) repository.ProfileRepository {
	return &profileRepository{
		collection: client.Database(dbName).Collection("profiles"),
	}
}

func (r *profileRepository) GetProfile(ctx context.Context, botID string) (*models.Profile, error) {
	var profile models.Profile
	err := r.collection.FindOne(ctx, bson.M{"botId": botID}).Decode(&profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *profileRepository) GetAllProfiles(ctx context.Context) ([]*models.Profile, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var profiles []*models.Profile
	if err := cursor.All(ctx, &profiles); err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r *profileRepository) UpsertProfile(ctx context.Context, profile *models.Profile) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"botId": profile.BotID}
	update := bson.M{"$set": profile}
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}
