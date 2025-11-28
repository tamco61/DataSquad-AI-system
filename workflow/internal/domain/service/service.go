package service

import (
	"context"
	"workflow/internal/model"
	"workflow/internal/repository"
)

type Service struct {
	userRepo repository.User
	taskRepo repository.Task
}

func NewService(user repository.User, task repository.Task) Service {
	return Service{
		userRepo: user,
		taskRepo: task,
	}
}

func (s Service) GetUser(ctx context.Context, userID string) (*model.UserDTO, error) {
	return s.userRepo.GetUser(ctx, userID)
}

func (s Service) CreateUser(ctx context.Context, u *model.UserDTO) error {
	return s.userRepo.CreateUser(ctx, u)
}

func (s Service) CreateTask(ctx context.Context, t *model.TaskDTO) error {
	return s.taskRepo.CreateTask(ctx, t)
}

func (s Service) GetTask(ctx context.Context, taskID string) (*model.TaskDTO, error) {
	return s.taskRepo.GetTask(ctx, taskID)
}

func (s Service) GetTasksFiltering(ctx context.Context, f model.TaskFilter) (*[]model.TaskDTO, error) {
	return s.taskRepo.GetTasksFiltering(ctx, f)
}
