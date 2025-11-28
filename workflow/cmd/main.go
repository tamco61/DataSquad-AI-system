package main

import (
	"context"
	"log"
	"os/signal"

	"syscall"
	"workflow/cmd/internal"
	"workflow/internal/api/http"
	"workflow/internal/config"
	"workflow/internal/domain/service"
	"workflow/internal/repository/postgres"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	config := config.LoadConfig()

	userDB, err := internal.NewPostgreSQL(config.UsersDBURL)
	if err != nil {
		log.Fatalf("Could not initialize Database connection %s", err)
	}
	defer userDB.Close()
	tasksDB, err := internal.NewPostgreSQL(config.TasksDBURL)
	if err != nil {
		log.Fatalf("Could not initialize Database connection %s", err)
	}
	defer tasksDB.Close()

	userRepo := postgres.NewUserRepo(userDB)
	taskRepo := postgres.NewTaskRepo(tasksDB)

	service := service.NewService(userRepo, taskRepo)

	http.RunService(ctx, config.Addr, service)

}
