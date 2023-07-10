package repositories

import (
	"focusic-api/config"
	"focusic-api/database"
	"focusic-api/models/domains"
)

type userRepository struct{}

func NewUserRepository() IBaseRepository[domains.User] {
	return &userRepository{}
}

func (userRepo userRepository) Create(entity domains.User) (domains.User, error) {
	var newUser domains.User
	appValidator := config.AppValidator
	err := appValidator.Struct(entity)
	if err != nil {
		return newUser, err
	}
	database.DB.Create(&entity)
	newUser = entity
	return newUser, nil
}
