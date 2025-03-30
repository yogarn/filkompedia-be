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
	books.Patch("/", r.middleware.Authenticate, r.middleware.Authorize([]int{1}), r.EditBook)
	books.Delete("/:id", r.middleware.Authenticate, r.middleware.Authorize([]int{1}), r.DeleteBook)
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
	carts.Use(r.middleware.Authenticate)

	carts.Get("/user/:userId", r.middleware.Authorize([]int{1}), r.GetUserCartAdmin)
	carts.Get("/user", r.GetUserCart)
	carts.Get("/:cartId", r.GetCart)
	carts.Post("/", r.AddToCart)
	carts.Patch("/", r.EditCart)
	carts.Delete("/:cartId", r.RemoveFromCart)
}

func mountCheckout(routerGroup fiber.Router, r *Rest) {
	checkouts := routerGroup.Group("/checkouts")
	checkouts.Use(r.middleware.Authenticate)

	checkouts.Get("/user/:userId", r.middleware.Authorize([]int{1}), r.GetUserCheckoutsAdmin)
	checkouts.Get("/user", r.GetUserCheckouts)
	checkouts.Get("/:checkoutId", r.GetCheckoutCarts)
	checkouts.Post("/", r.Checkout)
}

func mountPayment(routerGroup fiber.Router, r *Rest) {
	payments := routerGroup.Group("/payments")

	payments.Post("/webhook", r.HandleMidtransWebhook)
	payments.Get("/:id", r.middleware.Authenticate, r.middleware.Authorize([]int{1}), r.GetPayement)
	payments.Get("/book/:id", r.middleware.Authenticate, r.CheckUserBookPurchase)
	payments.Get("/", r.middleware.Authenticate, r.middleware.Authorize([]int{1}), r.GetPayments)
	payments.Get("/checkout/:id", r.middleware.Authenticate, r.GetPayementByCheckout)
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
	mountPayment(routerGroup, r)
}

func (r *Rest) Start(port string) error {
	return r.router.Listen(port)
}
