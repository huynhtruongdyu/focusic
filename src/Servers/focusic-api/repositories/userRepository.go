package repositories

import (
	"focusic-api/config"
	"focusic-api/database"
	"focusic-api/models/domains"
	"os"
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

func (userRepo userRepository) FindAll() []domains.User {
	var users = make([]domains.User, 0)
	query, err := os.ReadFile("resources/dbScripts/get_all_users.sql")
	if err != nil {
		return users
	}
	database.DB.Raw(string(query)).Scan(&users)
	return users
}
