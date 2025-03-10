package rest

import (
	"net/http"

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

	token, err := r.service.AuthService.Login(loginReq)
	if err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", token)
	return nil
}
