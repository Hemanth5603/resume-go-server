package service

import (
	"github.com/Hemanth5603/resume-go-server/internal/model"
	"github.com/Hemanth5603/resume-go-server/internal/repository"
	"github.com/Hemanth5603/resume-go-server/internal/utils"
)

// UserService defines the interface for user-related business logic.
type UserService interface {
	CreateUser(req *model.CreateUserRequest) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	UpdateUser(id string, user *model.User) (*model.User, error)
	DeleteUser(id string) error
	ListUsers(page, limit int64) ([]*model.User, error)
	CountUsers() (int64, error)
}

// userServiceImpl is the concrete implementation of the UserService.
type userServiceImpl struct {
	userRepo *repository.UserRepositoryMongo
}

// NewUserService creates a new UserService with the given user repository.
func NewUserService(userRepo *repository.UserRepositoryMongo) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

// CreateUser handles the business logic for creating a new user.
func (s *userServiceImpl) CreateUser(req *model.CreateUserRequest) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	return s.userRepo.CreateUser(user)
}

// GetUserByID retrieves a user by ID
func (s *userServiceImpl) GetUserByID(id string) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}

// GetUserByEmail retrieves a user by email
func (s *userServiceImpl) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

// GetUserByUsername retrieves a user by username
func (s *userServiceImpl) GetUserByUsername(username string) (*model.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

// UpdateUser updates a user
func (s *userServiceImpl) UpdateUser(id string, user *model.User) (*model.User, error) {
	return s.userRepo.UpdateUser(id, user)
}

// DeleteUser deletes a user
func (s *userServiceImpl) DeleteUser(id string) error {
	return s.userRepo.DeleteUser(id)
}

// ListUsers retrieves users with pagination
func (s *userServiceImpl) ListUsers(page, limit int64) ([]*model.User, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	return s.userRepo.ListUsers(page, limit)
}

// CountUsers returns total user count
func (s *userServiceImpl) CountUsers() (int64, error) {
	return s.userRepo.CountUsers()
}
