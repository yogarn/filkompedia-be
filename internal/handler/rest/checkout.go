package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/model"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func (r *Rest) GetUserCheckouts(ctx *fiber.Ctx) error {
	param := ctx.Params("userId")
	userId, err := uuid.Parse(param)
	if err != nil {
		return err
	}

	checkouts, err := r.service.CheckoutService.GetUserCheckouts(userId)
	if err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", checkouts)
	return nil
}

func (r *Rest) GetCheckoutCarts(ctx *fiber.Ctx) error {
	param := ctx.Params("checkoutId")
	checkoutId, err := uuid.Parse(param)
	if err != nil {
		return err
	}

	carts, err := r.service.CheckoutService.GetCheckoutCarts(checkoutId)
	if err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", carts)
	return nil
}

func (r *Rest) Checkout(ctx *fiber.Ctx) error {
	var request model.CheckoutRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	if err := r.service.CheckoutService.Checkout(request); err != nil {
		return err
	}

	//handle payments

	response.Success(ctx, http.StatusOK, "success", nil) //payment details

	return nil
}
