package handlers

import (
	"book_store/models"
	"log"

	"github.com/labstack/echo/v4"
)

type BookHandlers struct {
	Repo models.BooksQuerier
}

func (h *BookHandlers) Create(c echo.Context) error {
	b := models.Book{}
	err := c.Bind(&b)
	if err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}

	return nil
}
