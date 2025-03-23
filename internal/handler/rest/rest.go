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
	users.Get("/me", r.middleware.Authenticate, r.GetMe)
	users.Get("/:userId", r.GetUserProfile)
}

func mountBook(routerGroup fiber.Router, r *Rest) {
	books := routerGroup.Group("/books")
	books.Get("/", r.middleware.Authenticate, r.SearchBooks)
	books.Get("/:id", r.middleware.Authenticate, r.GetBook)
	books.Post("/", r.middleware.Authenticate, r.middleware.Authorize([]int{1}), r.CreateBook)
}

func mountComment(routerGroup fiber.Router, r *Rest) {
	comments := routerGroup.Group("/comments")
	comments.Use(r.middleware.Authenticate)

	comments.Get("/:id", r.GetComment)
	comments.Get("/book/:bookId", r.GetCommentByBook)
	comments.Post("/", r.CreateComment)
	comments.Put("/book/:bookId/comment/:id", r.UpdateComment)
	comments.Delete("/:id", r.DeleteComment)
}

func mountCart(routerGroup fiber.Router, r *Rest) {
	carts := routerGroup.Group("/carts")
	carts.Get("/user/:userId", r.middleware.Authenticate, r.GetUserCart)
	carts.Get("/:cartId", r.middleware.Authenticate, r.GetCart)
	carts.Post("/", r.middleware.Authenticate, r.AddToCart)
	carts.Delete("/:cartId", r.middleware.Authenticate, r.RemoveFromCart)
}

func mountCheckout(routerGroup fiber.Router, r *Rest) {
	checkouts := routerGroup.Get("/checkouts")
	checkouts.Get("/user/:userId", r.middleware.Authenticate, r.GetUserCheckouts)
	checkouts.Get("/:checkoutId", r.middleware.Authenticate, r.GetCheckoutCarts)
	checkouts.Post("/", r.middleware.Authenticate, r.Checkout)
}

func (r *Rest) RegisterRoutes() {
	routerGroup := r.router.Group("/api/v1")

	routerGroup.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	mountUser(routerGroup, r)
	mountAuth(routerGroup, r)
	mountBook(routerGroup, r)
	mountComment(routerGroup, r)
	mountCart(routerGroup, r)
	mountCheckout(routerGroup, r)
}

func (r *Rest) Start(port string) error {
	return r.router.Listen(port)
}
