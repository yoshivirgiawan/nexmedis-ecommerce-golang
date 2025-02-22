package user

import (
	"crypto/rand"
	"ecommerce/helper"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID string, fileLocation string) (User, error)
	GetUsersWithFilters(take int, skip int, search string) ([]User, error)
	GetUserByID(ID string) (User, error)
	CreateUser(input CreateUserInput) (User, error)
	UpdateUser(inputID GetUserDetailInput, input CreateUserInput) (User, error)
	DeleteUser(ID string) error
}

type service struct {
	repository Repository
}

func NewService() *service {
	repository := NewRepository()
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.PhoneNumber = input.PhoneNumber
	password, _ := GenerateRandomPassword(10)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	subject := "Welcome to Our Service!"
	htmlBody := fmt.Sprintf(`
		<html>
		<body>
			<h2>Welcome %s!</h2>
			<p>Thank you for registering with us. Your account has been created successfully.</p>
			<p>Your temporary password is: <strong>%s</strong></p>
		</body>
		</html>`, newUser.Name, password)

	err = helper.SendEmail(newUser.Email, subject, htmlBody)

	if err != nil {
		return newUser, errors.New("Failed to send email")
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == "" {
		return true, nil
	}

	return false, nil
}

func (s *service) IsPhoneNumberAvailable(input CheckPhoneNumberInput) (bool, error) {
	phone_number := input.PhoneNumber

	user, err := s.repository.FindByPhoneNumber(phone_number)
	if err != nil {
		return false, err
	}

	if user.ID == "" {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID string, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) GetUsersWithFilters(take int, skip int, search string) ([]User, error) {
	return s.repository.FindAllWithFilters(take, skip, search)
}

func (s *service) CreateUser(input CreateUserInput) (User, error) {
	duplicateUser, _ := s.repository.FindByEmail(input.Email)
	if duplicateUser.ID != "" {
		return duplicateUser, errors.New("Email already exists")
	}

	duplicateUser, _ = s.repository.FindByPhoneNumber(input.PhoneNumber)
	if duplicateUser.ID != "" {
		return duplicateUser, errors.New("Phone number already exists")
	}

	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.PhoneNumber = input.PhoneNumber
	user.Role = "user"

	password, _ := GenerateRandomPassword(10)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) UpdateUser(inputID GetUserDetailInput, input CreateUserInput) (User, error) {
	user, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return user, err
	}

	duplicateUser, _ := s.repository.FindByEmail(input.Email)
	if duplicateUser.ID != "" && duplicateUser.ID != inputID.ID {
		return duplicateUser, errors.New("Email already exists")
	}

	duplicateUser, _ = s.repository.FindByPhoneNumber(input.PhoneNumber)
	if duplicateUser.ID != "" && duplicateUser.ID != inputID.ID {
		return duplicateUser, errors.New("Phone number already exists")
	}

	user.Name = input.Name
	user.Email = input.Email
	user.PhoneNumber = input.PhoneNumber

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) DeleteUser(ID string) error {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return err
	}

	if user.ID == "" {
		return errors.New("No user found on with that ID")
	}

	_, err = s.repository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}

func GenerateRandomPassword(length int) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	var password strings.Builder

	for i := 0; i < length; i++ {
		// Pilih karakter acak dari string chars
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		password.WriteByte(chars[index.Int64()])
	}

	return password.String(), nil
}
