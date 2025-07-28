package service

import (
	"github.com/Hemanth5603/resume-go-server/internal/model"
	"github.com/Hemanth5603/resume-go-server/internal/repository"
	"github.com/Hemanth5603/resume-go-server/internal/utils"
)

// UserService defines the interface for user-related business logic.
type UserService interface {
	CreateUser(req *model.CreateUserRequest) (*model.User, error)
}

// userServiceImpl is the concrete implementation of the UserService.
type userServiceImpl struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new UserService with the given user repository.
func NewUserService(userRepo *repository.UserRepository) UserService {
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
