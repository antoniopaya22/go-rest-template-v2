package repository

import (
	"ai4am.com/go-restapi/internal/config"
	"ai4am.com/go-restapi/models"
	"strconv"
)

type UserRepository struct{}
var userRepository *UserRepository

// Singleton Desing Pattern
func GetUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{}
	}
	return userRepository
}

func (r *UserRepository) GetUserById(id string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &user, []string{})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.Username = username
	_, err := First(&where, &user, []string{})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	err := Find(&models.User{}, &users, []string{}, "id asc")
	return &users, err
}

func (r *UserRepository) QueryUsers(q *models.User) (*[]models.User, error) {
	var users []models.User
	err := Find(&q, &users, []string{}, "id asc")
	return &users, err
}

func (r *UserRepository) AddUser(user *models.User) error {
	err := Create(&user)
	err = Save(&user)
	return err
}
func (r *UserRepository) UpdateUser(user *models.User) error { return config.GetDB().Save(&user).Error }

func (r *UserRepository) DeleteUser(user *models.User) error { return config.GetDB().Unscoped().Delete(&user).Error }

