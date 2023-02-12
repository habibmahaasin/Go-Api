package service

import (
	"errors"
	"gop-api/modules/user/models"
	"gop-api/modules/user/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(inputLogin models.InputLogin) (models.User, error)
	AddUser(addUser models.AddUser) error
	UpdateUser(id string, inputUpdate models.User) (models.User, error)
	DeleteUser(id string) (models.User, error)
	GetUser() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) GetUser() ([]models.User, error) {
	return s.repository.GetUser()
}

func (s *userService) GetUserByEmail(email string) (models.User, error) {
	return s.repository.GetUserByEmail(email)
}

func (s *userService) GetUserById(id string) (models.User, error) {
	return s.repository.GetUserById(id)
}

func (s *userService) Login(inputLogin models.InputLogin) (models.User, error) {
	email := inputLogin.Email
	password := inputLogin.Password

	findUser, _ := s.repository.GetUserByEmail(email)
	if findUser.User_id == "" {
		return findUser, errors.New("Email Doesnt Exist")
	}

	err := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(password))
	if err != nil {
		return findUser, errors.New("Invalid Email or Password")
	}

	return findUser, nil
}

func (s *userService) AddUser(addUser models.AddUser) error {
	password := []byte(addUser.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	addUser.Password = string(hashedPassword)
	return s.repository.AddUser(addUser)
}

func (s *userService) UpdateUser(id string, inputUpdate models.User) (models.User, error) {
	password := []byte(inputUpdate.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	inputUpdate.Password = string(hashedPassword)

	return s.repository.UpdateUser(id, inputUpdate)
}

func (s *userService) DeleteUser(id string) (models.User, error) {
	return s.repository.DeleteUser(id)
}
