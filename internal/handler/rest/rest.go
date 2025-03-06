package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/filkompedia-be/internal/service"
)

type Rest struct {
	router  *fiber.App
	service *service.Service
}

func NewRest(router *fiber.App, service *service.Service) *Rest {
	return &Rest{
		router:  router,
		service: service,
	}
}

func mountUser(routerGroup fiber.Router, r *Rest) {
	users := routerGroup.Group("/users")
	users.Get("/", r.GetAllUserProfile)
	users.Get("/:userId", r.GetUserProfile)
}

func (r *Rest) RegisterRoutes() {
	routerGroup := r.router.Group("/api/v1")

	routerGroup.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	mountUser(routerGroup, r)
}

func (r *Rest) Start(port string) error {
	return r.router.Listen(port)
}
