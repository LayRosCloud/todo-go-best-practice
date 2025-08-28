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
)

func main() {
	const Port = 8080
	
	// Загрузка конфигурации
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error with config: %v", err)
	}

	// Подключение к БД
	db, err := databases.NewPostgresConnection(config)
	if err != nil {
		log.Fatalf("Error with database: %v", err)
	}
	
	// Гарантированное закрытие соединения
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Миграции
	if err = db.RunMigrations(config); err != nil {
		log.Fatalf("Error with migrations: %v", err)
		return
	}

	// Инициализация зависимостей
	taskMapper := mappers.TaskMapper{}
	userMapper := mappers.UserMapper{TaskMapper: taskMapper}

	// Убедитесь, что репозитории получают актуальное соединение
	taskRepo := repositories.CreateTaskRepository(db.DB)
	userRepo := repositories.CreateUserRepository(db.DB, taskRepo)

	userService := services.NewUserService(userRepo, userMapper)
	userHandler := handlers.NewUserHandler(*userService)

	paramsApi := api.NewApiSetup(userHandler)
	router := api.SetupRoutes(paramsApi)

	// Запуск сервера
	log.Printf("Server starting on port: %d\n", Port)
	if err := db.DB.Ping(); err != nil {
    	log.Fatalf("Database connection is dead before server start: %v", err)
	}

	log.Printf("Database connection is alive: %v", db.DB.Stats())
	if err := http.ListenAndServe(fmt.Sprintf(":%d", Port), router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}