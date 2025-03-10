package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/yogarn/filkompedia-be/internal/handler/rest"
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/internal/service"
	"github.com/yogarn/filkompedia-be/pkg/bcrypt"
	"github.com/yogarn/filkompedia-be/pkg/jwt"
)

type Config struct {
	DB  *sqlx.DB
	App *fiber.App
}

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func StartUp(config *Config) {
	bcrypt := bcrypt.Init()
	jwt := jwt.Init()

	repository := repository.NewRepository(config.DB)
	service := service.NewService(repository, bcrypt, jwt)

	rest := rest.NewRest(config.App, service)
	rest.RegisterRoutes()

	rest.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
