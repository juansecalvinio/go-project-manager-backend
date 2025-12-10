package services

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"

	"go-project-manager-backend/internal/domain/models"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
	List() ([]*models.User, error)
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Register(name, email, password string, role models.Role) (*models.User, error) {
	_, err := s.repository.GetByEmail(email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword := hashPassword(password)
	user := &models.User{
		ID:             generateID(),
		Name:           name,
		Email:          email,
		PasswordHashed: hashedPassword,
		Role:           role,
	}

	err = s.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.repository.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !checkPassword(password, user.PasswordHashed) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.repository.GetByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.repository.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repository.Delete(id)
}

func (s *UserService) ListUsers() ([]*models.User, error) {
	return s.repository.List()
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}

func checkPassword(password, hash string) bool {
	return hashPassword(password) == hash
}

func generateID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
