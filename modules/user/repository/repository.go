package repository

import (
	"gop-api/modules/user/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser() ([]models.User, error)
	AddUser(addUser models.AddUser) error
	GetUserByEmail(email string) (models.User, error)
	GetUserById(id string) (models.User, error)
	UpdateUser(id string, inputUpdate models.User) (models.User, error)
	DeleteUser(id string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUser() ([]models.User, error) {
	var users []models.User
	err := r.db.Raw("SELECT user_id, name, email FROM users").Scan(&users).Error
	return users, err
}

func (r *userRepository) GetUserById(id string) (models.User, error) {
	var users models.User
	err := r.db.Raw("SELECT user_id, name, email FROM users WHERE user_id=?", id).Scan(&users).Error
	return users, err
}

func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}

func (r *userRepository) AddUser(addUser models.AddUser) error {
	err := r.db.Raw("INSERT INTO users (user_id, name, email, password, date_created, date_updated) VALUES(uuid(),?,?,?,?,?)", addUser.Name, addUser.Email, addUser.Password, time.Now(), time.Now()).Scan(&addUser).Error
	return err
}

func (r *userRepository) UpdateUser(id string, inputUpdate models.User) (models.User, error) {
	var user models.User
	err := r.db.Raw("UPDATE users SET name=?, email=?, password=?, date_updated=? WHERE user_id=?", inputUpdate.Name, inputUpdate.Email, inputUpdate.Password, time.Now(), id).Scan(&user).Error
	return user, err
}

func (r *userRepository) DeleteUser(id string) (models.User, error) {
	var user models.User
	err := r.db.Raw("DELETE FROM users WHERE user_id=?", id).Scan(&user).Error
	return user, err
}
