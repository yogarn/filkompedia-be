package rest

import "github.com/gofiber/fiber/v2"

type Rest struct {
	router *fiber.App
}

func NewRest(router *fiber.App) *Rest {
	return &Rest{
		router: router,
	}
}

func (r *Rest) RegisterRoutes() {
	routerGroup := r.router.Group("/api/v1")

	routerGroup.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
}

func (r *Rest) Start(port string) error {
	return r.router.Listen(port)
}
