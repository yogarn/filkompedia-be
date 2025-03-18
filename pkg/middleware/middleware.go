package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/filkompedia-be/internal/service"
	"github.com/yogarn/filkompedia-be/pkg/jwt"
)

type IMiddleware interface {
	Authenticate(ctx *fiber.Ctx) error
	Authorize(roles []int) fiber.Handler
}

type middleware struct {
	jwtAuth jwt.IJwt
	service *service.Service
}

func Init(jwtAuth jwt.IJwt, service *service.Service) IMiddleware {
	return &middleware{
		jwtAuth: jwtAuth,
		service: service,
	}
}
