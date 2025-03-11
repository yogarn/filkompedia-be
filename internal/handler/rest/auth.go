package rest

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/filkompedia-be/model"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func (r *Rest) Register(ctx *fiber.Ctx) (err error) {
	registerReq := &model.RegisterReq{}
	if err := ctx.BodyParser(registerReq); err != nil {
		return err
	}

	user, err := r.service.AuthService.Register(registerReq)
	if err != nil {
		return err
	}

	response.Success(ctx, http.StatusCreated, "success", user)
	return nil
}

func (r *Rest) Login(ctx *fiber.Ctx) (err error) {
	loginReq := &model.LoginReq{}
	if err := ctx.BodyParser(loginReq); err != nil {
		return err
	}

	loginRes, err := r.service.AuthService.Login(loginReq)
	if err != nil {
		return err
	}

	expiresIn, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME"))
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    loginRes.JwtToken,
		Expires:  time.Now().Add(time.Duration(expiresIn) * time.Second),
		HTTPOnly: true,
		Secure:   true,
		Path:     "/",
	})

	response.Success(ctx, http.StatusOK, "success", nil)
	return nil
}
