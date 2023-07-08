package repositories

import (
	"context"
	"errors"
	"fmt"
	"focusic-api/database"
	"focusic-api/models/domains"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var userCollection *mongo.Collection = database.GetCollection(database.MongoDbInstance, "users")

func GetUserRepository() IBaseRepository[domains.User] {
	return &userRepository{}
}

type userRepository struct{}

func (repo userRepository) Create(entity domains.User) (domains.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	utcNow := time.Now().UTC()
	entity.CreatedAt = utcNow
	entity.UpdatedAt = utcNow
	userResult, err := userCollection.InsertOne(ctx, entity)
	if err != nil {
		return domains.User{}, errors.New(err.Error())
	}
	objectID := userResult.InsertedID.(primitive.ObjectID)
	entity.ID = objectID

	fmt.Println(userResult)
	return entity, nil
}

func (repo userRepository) FindAll() ([]domains.User, error) {
	users := make([]domains.User, 0)
	return users, nil
}
