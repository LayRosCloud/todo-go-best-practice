package main

import (
	"flag"
	databases "leafall/todo-service/pkg/database"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	var (
        conn      = flag.String("conn", os.Getenv("database_conn"), "Database connection string")
        migrationsPath = flag.String("migrations", "./internal/database/migrations", "Path to migrations")
        command      = flag.String("command", "up", "Migration command: up, down, force, version")
        version      = flag.Int("version", 0, "Version for force migration")
    )
    flag.Parse()

	db, err := databases.NewPostgresConnectionString(*conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db.DB.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://" + *migrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	switch *command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations applied successfully")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations reverted successfully")
	case "force":
		if err := m.Force(*version); err != nil {
			log.Fatal(err)
		}
		log.Printf("Migration forced to version %d", *version)
	case "version":
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Current version: %d, dirty: %t", version, dirty)
	default:
		log.Fatalf("Unknown command: %s", *command)
	}
}