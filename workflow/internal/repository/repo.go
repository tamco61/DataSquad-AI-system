package repository

import (
	"context"
	"workflow/internal/model"
)

type User interface {
	GetUser(context.Context, string) (*model.UserDTO, error)
	CreateUser(ctx context.Context, u *model.UserDTO) error
	GetUsers(ctx context.Context) (*[]model.UserDTO, error)
}

type Task interface {
	GetTask(context.Context, string) (*model.TaskDTO, error)
	GetTasksFiltering(context.Context, model.TaskFilter) (*[]model.TaskDTO, error)
	CreateTask(ctx context.Context, t *model.TaskDTO) error
}
