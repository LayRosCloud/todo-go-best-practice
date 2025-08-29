package main

import (
	"fmt"
	"leafall/todo-service/internal/api"
	"leafall/todo-service/internal/config"
	"leafall/todo-service/internal/handlers"
	"leafall/todo-service/internal/mappers"
	"leafall/todo-service/internal/repositories"
	"leafall/todo-service/internal/services"
	databases "leafall/todo-service/pkg/database"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	const Port = 8080
	db, config := LoadDatabaseAndConfig()
	defer db.Close()
	router := ConfigureRouter(db, config)
	
	ListenAndServe(router, Port, *db)
}

func LoadDatabaseAndConfig() (*databases.Database, *config.Config) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error with config: %v", err)
	}

	db, err := databases.NewPostgresConnection(config)
	if err != nil {
		log.Fatalf("Error with database: %v", err)
	}
	
	if err = db.RunMigrations(config); err != nil {
		log.Fatalf("Error with migrations: %v", err)
	}
	return db, config
}

func ConfigureRouter(db *databases.Database, config *config.Config) *chi.Mux {
	taskMapper := mappers.TaskMapper{}
	userMapper := mappers.NewUserMapper(&taskMapper)

	taskRepo := repositories.CreateTaskRepository(db.DB)
	tokenRepo := repositories.CreateTokenRepository(db.DB)
	userRepo := repositories.CreateUserRepository(db.DB, taskRepo)

	accessTokenService := services.NewTokenService(config.AccessSecret, time.Minute * 15)
	refreshTokenService := services.NewTokenService(config.RefreshSecret, time.Hour * 24 * 30)
	userService := services.NewUserService(userRepo, userMapper, accessTokenService, refreshTokenService, tokenRepo)
	taskService := services.NewTaskService(taskRepo, taskMapper)

	userHandler := handlers.NewUserHandler(userService)
	taskHandler := handlers.NewTaskHandler(taskService)

	paramsApi := api.NewApiSetup(userHandler, taskHandler, userService)
	router := api.SetupRoutes(paramsApi)
	return router
}

func ListenAndServe(router *chi.Mux, port int, db databases.Database) {
	log.Printf("Server starting on port: %d\n", port)
	if err := db.DB.Ping(); err != nil {
    	log.Fatalf("Database connection is dead before server start: %v", err)
	}

	log.Printf("Database connection is alive: %v", db.DB.Stats())
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}