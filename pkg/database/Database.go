package databases

import (
	"fmt"
	"leafall/todo-service/internal/config"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
    DB *sqlx.DB
}


func (d *Database) ToConnectionString(cfg *config.Config) string {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.PortDb, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	return connStr
}

func NewPostgresConnection(cfg *config.Config) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.PortDb, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	return NewPostgresConnectionString(connStr)
}

func NewPostgresConnectionString(connection string) (*Database, error) {
	db, err := sqlx.Connect("postgres", connection)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database")
	return &Database{DB: db}, nil
}

func (d *Database) RunMigrations(conf *config.Config) error {
	log.Println("Running migrations with temporary connection...")
    
    tempDB, err := sqlx.Connect("postgres", d.ToConnectionString(conf))
    if err != nil {
        return fmt.Errorf("failed to create temporary database connection: %w", err)
    }
    defer tempDB.Close()

    driver, err := postgres.WithInstance(tempDB.DB, &postgres.Config{})
    if err != nil {
        return fmt.Errorf("failed to create migration driver: %w", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        fmt.Sprintf("file://%s", conf.MigrationPath),
        "postgres", 
        driver,
    )
    if err != nil {
        return fmt.Errorf("failed to create migration instance: %w", err)
    }
    defer m.Close()

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("failed to run migrations: %w", err)
    }

    log.Println("Migrations applied successfully")
    return nil
}

func (d *Database) Close() error {
	log.Println("Closing database connection...")
    return d.DB.Close()
}