package service

import (
	"errors"
	"gop-api/modules/user/models"
	"gop-api/modules/user/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUser() ([]models.User, error)
	AddUser(addUser models.AddUser) error
	Login(inputLogin models.InputLogin) (models.User, error)
	GetUserById(id string) (models.User, error)
	UpdateUser(id string, inputUpdate models.User) (models.User, error)
	DeleteUser(id string) (models.User, error)
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

func (s *userService) GetUserById(id string) (models.User, error) {
	return s.repository.GetUserById(id)
}

func (s *userService) UpdateUser(id string, inputUpdate models.User) (models.User, error) {
	password := []byte(inputUpdate.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	inputUpdate.Password = string(hashedPassword)

	findUser, _ := s.repository.FindUserByEmail(inputUpdate.Email)
	if findUser.Email != "" {
		return findUser, errors.New("Email Already Exists")
	}

	return s.repository.UpdateUser(id, inputUpdate)
}

func (s *userService) AddUser(addUser models.AddUser) error {
	password := []byte(addUser.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	generateUUID := uuid.Must(uuid.NewRandom())
	findUser, err := s.repository.FindUserByEmail(addUser.Email)

	if findUser.Email != "" {
		return errors.New("Email Already Exists")
	}

	addUser.User_uuid = generateUUID
	addUser.Password = string(hashedPassword)
	return s.repository.AddUser(addUser)
}

func (s *userService) Login(inputLogin models.InputLogin) (models.User, error) {
	email := inputLogin.Email
	password := inputLogin.Password

	findUser, _ := s.repository.FindUserByEmail(email)
	if findUser.User_uuid == "" {
		return findUser, errors.New("User Doesnt Exist")
	}

	err := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(password))
	if err != nil {
		return findUser, errors.New("Invalid Email or Password")
	}

	return findUser, nil
}

func (s *userService) DeleteUser(id string) (models.User, error) {
	return s.repository.DeleteUser(id)
}
