package handler

import (
	// "fmt"

	"altech-omega-api/domain"
	helper_http "altech-omega-api/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type BookHandler struct {
	usecase domain.BookUsecase
}

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

func (h *BookHandler) GetAllHandler(c echo.Context) error {
	data, err := h.usecase.GetAll()

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all book")

	return resp
}

func (h *BookHandler) GetByIDHandler(c echo.Context) error {
	id := c.Param("id")
	// id = fmt.Sprintf("%s")
	// num, err := strconv.Atoi(id)

	// if err != nil {
	// 	panic(err)
	// }

	data, err := h.usecase.GetByID(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, data, "success get by id")
	return resp
}


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

func (h *BookHandler) DeleteHandler(c echo.Context) error {
	id := c.Param("id")

	err := h.usecase.Delete(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, nil, "success delete book")
	return resp
}
