package repositories

import (
	"context"

	"medods-test-task/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type SessionRepo struct {
	db *mongo.Collection
}

func NewSessionRepo(db *mongo.Database) *SessionRepo {
	return &SessionRepo{db: db.Collection("sessions")}
}

func (r *SessionRepo) CreateSession(ctx context.Context, session entities.Session) error {
	_, err := r.db.InsertOne(ctx, session)
  	return err
}

func (r *SessionRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (entities.Session, error) {
	var session entities.Session
	err := r.db.FindOne(ctx, bson.M{
		"refreshToken": refreshToken,
	}).Decode(&session)

	return session, err
}

func (r *SessionRepo) DeleteSession(ctx context.Context, session entities.Session) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": session.ID})

	return err
}