package domain

import (
	"context"
	"workflow/internal/model"
)

type Service interface {
	GetTask(ctx context.Context, userID string) (*model.TaskDTO, error)
	GetUser(ctx context.Context, userID string) (*model.UserDTO, error)
	CreateTask(ctx context.Context, t *model.TaskDTO) error
	CreateUser(ctx context.Context, u *model.UserDTO) error
	GetTasksFiltering(context.Context, model.TaskFilter) (*[]model.TaskDTO, error)
}
