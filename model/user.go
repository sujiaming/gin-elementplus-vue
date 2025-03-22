package model

import (
	"wanworld/database"
)

type UserRepository struct{}

func (r *UserRepository) FindAll() ([]database.User, error) {
	var users []database.User
	result := database.DB.Find(&users)
	return users, result.Error
}

func (r *UserRepository) Create(user *database.User) error {
	result := database.DB.Create(user)
	return result.Error
}

func (r *UserRepository) FindByID(id string) (*database.User, error) {
	var user database.User
	result := database.DB.First(&user, id)
	return &user, result.Error
}

func (r *UserRepository) Update(user *database.User) error {
	result := database.DB.Save(user)
	return result.Error
}

func (r *UserRepository) Delete(id string) error {
	result := database.DB.Delete(&database.User{}, id)
	return result.Error
}
