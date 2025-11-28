package config

import "os"

type Config struct {
	UsersDBURL string
	TasksDBURL string
	Addr       string
}

func LoadConfig() *Config {
	port, _ := os.LookupEnv("WORKFLOW_PORT")
	usersDB, _ := os.LookupEnv("USERS_DB_URL")
	tasksDB, _ := os.LookupEnv("TASKS_DB_URL")
	return &Config{
		Addr:       port,
		TasksDBURL: tasksDB,
		UsersDBURL: usersDB,
	}
}
