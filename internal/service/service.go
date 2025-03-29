package service

import (
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/pkg/bcrypt"
	"github.com/yogarn/filkompedia-be/pkg/jwt"
	"github.com/yogarn/filkompedia-be/pkg/midtrans"
	"github.com/yogarn/filkompedia-be/pkg/smtp"
)

type Service struct {
	UserService     IUserService
	AuthService     IAuthService
	BookService     IBookService
	CartService     ICartService
	CommentService  ICommentService
	CheckoutService ICheckoutService
	PaymentService  IPaymentService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.IBcrypt, jwt jwt.IJwt, smtp *smtp.SMTPClient, midtrans midtrans.IMidtrans) *Service {
	return &Service{
		UserService:     NewUserService(repository.UserRepository),
		AuthService:     NewAuthService(repository.AuthRepository, repository.UserRepository, bcrypt, jwt, smtp),
		BookService:     NewBookService(repository.BookRepository, repository.CartRepository),
		CartService:     NewCartService(repository.CartRepository, repository.UserRepository, repository.BookRepository),
		CommentService:  NewCommentService(repository.CommentRepository, repository.UserRepository),
		CheckoutService: NewCheckoutService(repository.CheckoutRepository, repository.CartRepository, repository.BookRepository),
		PaymentService:  NewPaymentService(repository.PaymentRepository, midtrans, repository.UserRepository, repository.BookRepository),
	}
}
