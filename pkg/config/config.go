package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/yogarn/filkompedia-be/internal/handler/rest"
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/internal/service"
	"github.com/yogarn/filkompedia-be/pkg/bcrypt"
	"github.com/yogarn/filkompedia-be/pkg/jwt"
	"github.com/yogarn/filkompedia-be/pkg/middleware"
	"github.com/yogarn/filkompedia-be/pkg/smtp"
)

type Config struct {
	DB    *sqlx.DB
	Redis *redis.Client
	App   *fiber.App
}

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func StartUp(config *Config) {
	bcrypt := bcrypt.Init()
	jwt := jwt.Init()
	smtp := smtp.LoadSMTPCredentials()

	repository := repository.NewRepository(config.DB, config.Redis)
	service := service.NewService(repository, bcrypt, jwt, smtp)

	middleware := middleware.Init(jwt, service)

	rest := rest.NewRest(config.App, service, middleware)
	rest.RegisterRoutes()

	rest.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
