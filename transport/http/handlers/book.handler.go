package handler

import (
	"altech-omega-api/domain"
	helper_http "altech-omega-api/transport/http/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	usecase domain.BookUsecase
}

// BookRoute registers routes for book operations
// @Summary Book Routes
// @Description Routes related to book management
// @ID book-routes
func BookRoute(e *echo.Echo, uc domain.BookUsecase) {
	handler := BookHandler{
		usecase: uc,
	}
	e.GET("/book/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/book")
	})

	e.GET("/book", handler.GetAllHandler)
	e.GET("/book/:id", handler.GetByIDHandler)
	e.POST("/book", handler.Create)
	e.PUT("/book/:id", handler.UpdateHandler)
	e.DELETE("/book/:id", handler.DeleteHandler)
}

// GetAllHandler handles the request to get all books
// @Summary Get all books
// @Description Retrieve a list of all books
// @ID get-all-books
// @Produce json
// @Success 200 {array} domain.Book
// @Failure 500 {object} helper_http.Response
// @Tags Book
// @Router /book [get]
func (h *BookHandler) GetAllHandler(c echo.Context) error {
	data, err := h.usecase.GetAll()

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all book")

	return resp
}

// GetByIDHandler handles the request to get a book by ID
// @Summary Get a book by ID
// @Description Retrieve a book by its unique ID
// @ID get-book-by-id
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} domain.Book
// @Failure 404 {object} helper_http.Response
// @Failure 500 {object} helper_http.Response
// @Tags Book
// @Router /book/{id} [get]
func (h *BookHandler) GetByIDHandler(c echo.Context) error {
	id := c.Param("id")

	data, err := h.usecase.GetByID(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}

// Create handles the request to create a new book
// @Summary Create a new book
// @Description Create a new book record
// @ID create-book
// @Accept json
// @Produce json
// @Param book body domain.Book true "Book data"
// @Success 201 {object} domain.Book
// @Failure 400 {object} helper_http.Response
// @Failure 422 {object} helper_http.Response
// @Tags Book
// @Router /book [post]
func (h *BookHandler) Create(c echo.Context) error {
	var data domain.Book

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create book")
}

// UpdateHandler handles the request to update a book by ID
// @Summary Update a book by ID
// @Description Update an existing book record by its ID
// @ID update-book-by-id
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body domain.Book true "Updated book data"
// @Success 200 {object} domain.Book
// @Failure 400 {object} helper_http.Response
// @Failure 404 {object} helper_http.Response
// @Failure 500 {object} helper_http.Response
// @Tags Book
// @Router /book/{id} [put]
func (h *BookHandler) UpdateHandler(c echo.Context) error {
	id := c.Param("id")

	var data domain.Book

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Update(id, &data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success update book")
	return resp
}

// DeleteHandler handles the request to delete a book by ID
// @Summary Delete a book by ID
// @Description Delete a book record by its ID
// @ID delete-book-by-id
// @Param id path string true "Book ID"
// @Success 200 {object} helper_http.Response
// @Failure 404 {object} helper_http.Response
// @Failure 500 {object} helper_http.Response
// @Tags Book
// @Router /book/{id} [delete]
func (h *BookHandler) DeleteHandler(c echo.Context) error {
	id := c.Param("id")

	err := h.usecase.Delete(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, nil, "success delete book")
	return resp
}
