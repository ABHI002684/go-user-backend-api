package main

import (
	"database/sql"
	"log"
	"os"

    "github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/ABHI002684/go-user-backend-api/db/sqlc"
	"github.com/ABHI002684/go-user-backend-api/internal/handler"
	"github.com/ABHI002684/go-user-backend-api/internal/logger"
	"github.com/ABHI002684/go-user-backend-api/internal/repository"
	"github.com/ABHI002684/go-user-backend-api/internal/routes"
	"github.com/ABHI002684/go-user-backend-api/internal/service"
)

func main() {

    if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Logger
	logg := logger.NewLogger()
	defer logg.Sync()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not set")
	}
	

	// Connect using database/sql (SQLC compatible)
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// SQLC Queries
	queries := sqlc.New(db)

	// App layers
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc, logg)

	// Fiber app
	app := fiber.New()
	routes.RegisterUserRoutes(app, h)

	log.Fatal(app.Listen(":3000"))
}
