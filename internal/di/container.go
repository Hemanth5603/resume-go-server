package di

import (
	"github.com/Hemanth5603/resume-go-server/configs"
	"github.com/Hemanth5603/resume-go-server/internal/database"
	"github.com/Hemanth5603/resume-go-server/internal/repository"
	"github.com/Hemanth5603/resume-go-server/internal/service"
)

type Container struct {
	Config      configs.Config
	UserService service.UserService
}

func NewContainer() (*Container, error) {
	cfg, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	return &Container{
		Config:      cfg,
		UserService: userService,
	}, nil
}
