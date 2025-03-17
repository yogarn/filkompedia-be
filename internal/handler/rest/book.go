package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/model"
	"github.com/yogarn/filkompedia-be/pkg/response"
)

func (r *Rest) GetBooks(ctx *fiber.Ctx) error {
	var bookReq model.BookReq
	if err := ctx.BodyParser(&bookReq); err != nil {
		return err
	}

	var books []entity.Book
	if err := r.service.BookService.GetBooks(&books, bookReq); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", books)
	return nil
}

func (r *Rest) SearchBooks(ctx *fiber.Ctx) error {
	var bookSearch model.BookSearch
	if err := ctx.BodyParser(&bookSearch); err != nil {
		return err
	}

	var books []entity.Book
	if err := r.service.BookService.SearchBooks(&books, bookSearch); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", books)
	return nil
}

func (r *Rest) CreateBook(ctx *fiber.Ctx) error {
	var create model.CreateBook
	if err := ctx.BodyParser(&create); err != nil {
		return err
	}

	if err := r.service.BookService.CreateBook(&create); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", nil)
	return nil
}
