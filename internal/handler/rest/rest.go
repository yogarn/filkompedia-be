package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/filkompedia-be/internal/service"
	"github.com/yogarn/filkompedia-be/pkg/middleware"
)

type Rest struct {
	router     *fiber.App
	service    *service.Service
	middleware middleware.IMiddleware
}

func NewRest(router *fiber.App, service *service.Service, middleware middleware.IMiddleware) *Rest {
	return &Rest{
		router:     router,
		service:    service,
		middleware: middleware,
	}
}

func mountAuth(routerGroup fiber.Router, r *Rest) {
	auths := routerGroup.Group("/auths")
	auths.Post("/register", r.Register)
	auths.Post("/login", r.Login)
	auths.Get("/sessions", r.middleware.Authenticate, r.GetSessions)
	auths.Post("/refresh", r.ExchangeToken)
	auths.Post("/send-otp", r.SendOtp)
	auths.Post("/verify-otp", r.VerifyOtp)
}

func mountUser(routerGroup fiber.Router, r *Rest) {
	users := routerGroup.Group("/users")
	users.Get("/", r.middleware.Authenticate, r.middleware.Authorize([]int{1}), r.GetAllUserProfile)
	users.Get("/:userId", r.GetUserProfile)
}

func mountBook(routerGroup fiber.Router, r *Rest) {
	books := routerGroup.Group("/books")
	books.Get("/", r.middleware.Authenticate, r.SearchBooks)
	books.Post("/", r.middleware.Authenticate, r.middleware.Authorize([]int{1}), r.CreateBook)
}

func mountCart(routerGroup fiber.Router, r *Rest) {
	carts := routerGroup.Group("/carts")
	carts.Get("/user/:userId", r.GetUserCart)
	carts.Get("/:cartId", r.GetCart)
	carts.Post("/", r.AddToCart)
	carts.Delete("/:cartId", r.RemoveFromCart)
}

func mountCheckout(routerGroup fiber.Router, r *Rest) {
	checkouts := routerGroup.Get("/checkouts")
	checkouts.Get("/user/:userId", r.GetUserCheckouts)
	checkouts.Get("/:checkoutId", r.GetCheckoutCarts)
	checkouts.Post("/", r.Checkout)
}

func (r *Rest) RegisterRoutes() {
	routerGroup := r.router.Group("/api/v1")

	routerGroup.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	mountUser(routerGroup, r)
	mountAuth(routerGroup, r)
	mountBook(routerGroup, r)
	mountCart(routerGroup, r)
	mountCheckout(routerGroup, r)
}

func (r *Rest) Start(port string) error {
	return r.router.Listen(port)
}
