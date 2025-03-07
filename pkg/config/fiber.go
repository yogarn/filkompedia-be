package config

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func StartFiber() *fiber.App {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: CustomErrorHandler,
		},
	)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept",
		AllowCredentials: false,
	}))

	return app
}

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	fmt.Print(err.Error())

	var errorRequest *response.ErrorResponse
	if errors.As(err, &errorRequest) {
		code = errorRequest.Code
		message = errorRequest.Message
	}

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		code = fiberError.Code
	}

	response.Error(ctx, code, message, err)

	return nil
}
