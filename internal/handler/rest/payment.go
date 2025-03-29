package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func (r *Rest) HandleMidtransWebhook(ctx *fiber.Ctx) error {
	var PaymentDetails map[string]any
	if err := ctx.BodyParser(&PaymentDetails); err != nil {
		return err
	}

	if err := r.service.PaymentService.UpdatePaymentStatus(PaymentDetails); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", nil)
	return nil
}

func (r *Rest) GetPayement(ctx *fiber.Ctx) error {
	paymentIdString := ctx.Params("id")
	paymentId, err := uuid.Parse(paymentIdString)
	if err != nil {
		return err
	}

	payment, err := r.service.PaymentService.GetPayment(paymentId)
	if err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", payment)
	return nil
}
