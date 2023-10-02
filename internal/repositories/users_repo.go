package repositories

import (
	"context"
	"log"
	

	"medods-test-task/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UsersRepo struct {
	db *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{db: db.Collection("users")}
}

func (r *UsersRepo) CreateUser(user entities.User) error {
	id, err := r.db.InsertOne(context.TODO(), user)
	log.Println(id)

  	return err
}

func (r *UsersRepo) GetByCredentials(ctx context.Context, email, password string) (entities.User, error) {
	var user entities.User
	err := r.db.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user) 

	return user, err
}

func (r *UsersRepo) UserExistsByEmail(ctx context.Context, email string) (bool) {
	var user entities.User
	err := r.db.FindOne(ctx, bson.M{"email": email}).Decode(&user) 

	return err == nil
}

func (r *UsersRepo) GetById(ctx context.Context, id primitive.ObjectID) (entities.User, error) {
	var user entities.User
	err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&user) 

	return user, err
}

