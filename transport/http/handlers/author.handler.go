package handler

import (
	"altech-omega-api/domain"
	helper_http "altech-omega-api/transport/http/helper"
	"net/http"
	"github.com/labstack/echo/v4"
)

// AuthorHandler defines handler methods for Author entity
type AuthorHandler struct {
	usecase domain.AuthorUsecase
}

// @Summary Author Routes
// @Description Routes related to author management
// @ID author-routes
func AuthorRoute(e *echo.Echo, uc domain.AuthorUsecase) {
	handler := AuthorHandler{
		usecase: uc,
	}
	e.GET("/author/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/author")
	})

	e.GET("/author", handler.GetAllHandler)
	e.GET("/author/:id", handler.GetByIDHandler)
	e.GET("/author/:id/books", handler.GetByIDWithBooksHandler)
	e.POST("/author", handler.Create)
	e.PUT("/author/:id", handler.UpdateHandler)
	e.DELETE("/author/:id", handler.DeleteHandler)
}

// GetAllHandler handles the request to get all authors
// @Summary Get all authors
// @Description Retrieve a list of all authors
// @ID get-all-authors
// @Produce json
// @0Success 200 {array} domain.Author
// @Failure 50 {object} helper_http.Response
// @Tags Author
// @Router /author [get]
func (h *AuthorHandler) GetAllHandler(c echo.Context) error {
	data, err := h.usecase.GetAll()

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all author")
	return resp
}

// GetByIDHandler handles the request to get an author by ID
// @Summary Get an author by ID
// @Description Retrieve an author by their unique ID
// @ID get-author-by-id
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} domain.Author
// @Failure 404 {object} helper_http.Response
// @Failure 500 {object} helper_http.Response
// @Tags Author
// @Router /author/{id} [get]
func (h *AuthorHandler) GetByIDHandler(c echo.Context) error {
	id := c.Param("id")

	data, err := h.usecase.GetByID(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}

// GetByIDWithBooksHandler handles the request to get an author by ID with their books
// @Summary Get an author with books by ID
// @Description Retrieve an author by their ID along with a list of their books
// @ID get-author-with-books-by-id
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} domain.Author
// @Failure 404 {object} helper_http.Response
// @Failure 500 {object} helper_http.Response
// @Tags Author
// @Router /author/{id}/books [get]
func (h *AuthorHandler) GetByIDWithBooksHandler(c echo.Context) error {
	id := c.Param("id")

	data, err := h.usecase.GetByIDWithBooks(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id with author's book")
	return resp
}

// Create handles the request to create a new author
// @Summary Create a new author
// @Description Create a new author record
// @ID create-author
// @Accept json
// @Produce json
// @Param author body domain.Author true "Author data"
// @Success 201 {object} domain.Author
// @Failure 400 {object} helper_http.Response
// @Failure 422 {object} helper_http.Response
// @Tags Author
// @Router /author [post]
func (h *AuthorHandler) Create(c echo.Context) error {
	var data domain.Author

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Create(&data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	return helper_http.SuccessResponse(c, data, "success create author")
}

// UpdateHandler handles the request to update an author by ID
// @Summary Update an author by ID
// @Description Update an existing author record by their ID
// @ID update-author-by-id
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Param author body domain.Author true "Updated author data"
// @Success 200 {object} domain.Author
// @Failure 400 {object} helper_http.Response
// @Failure 404 {object} helper_http.Response
// @Failure 500 {object} helper_http.Response
// @Tags Author
// @Router /author/{id} [put]
func (h *AuthorHandler) UpdateHandler(c echo.Context) error {
	id := c.Param("id")

	var data domain.Author

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.usecase.Update(id, &data)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success update author")
	return resp
}

// DeleteHandler handles the request to delete an author by ID
// @Summary Delete an author by ID
// @Description Delete an author record by their ID
// @ID delete-author-by-id
// @Param id path string true "Author ID"
// @Success 200 {object} helper_http.Response
// @Failure 404 {object} helper_http.Response
// @Failure 500 {object} helper_http.Response
// @Tags Author
// @Router /author/{id} [delete]
func (h *AuthorHandler) DeleteHandler(c echo.Context) error {
	id := c.Param("id")

	err := h.usecase.Delete(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, nil, "success delete author")
	return resp
}
