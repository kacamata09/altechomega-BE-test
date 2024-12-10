package handler

import (
	// "fmt"

	"altech-omega-api/domain"
	helper_http "altech-omega-api/transport/http/helper"
	"net/http"

	// "strconv"

	"github.com/labstack/echo"
)

type AuthorHandler struct {
	usecase domain.AuthorUsecase
}

func AuthorRoute(e *echo.Echo, uc domain.AuthorUsecase) {
	handler := AuthorHandler{
		usecase: uc,
	}
	e.GET("/author/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/author")
	})

	e.GET("/author", handler.GetAllHandler)
	e.GET("/author/:id", handler.GetByIDHandler)
	e.POST("/author", handler.Create)
	e.PUT("/author/:id", handler.UpdateHandler)
	e.DELETE("/author/:id", handler.DeleteHandler)
}

func (h *AuthorHandler) GetAllHandler(c echo.Context) error {
	data, err := h.usecase.GetAll()

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}
	resp := helper_http.SuccessResponse(c, data, "success get all author")

	return resp
}

func (h *AuthorHandler) GetByIDHandler(c echo.Context) error {
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

func (h *AuthorHandler) DeleteHandler(c echo.Context) error {
	id := c.Param("id")

	err := h.usecase.Delete(id)

	if err != nil {
		return helper_http.ErrorResponse(c, err)
	}

	resp := helper_http.SuccessResponse(c, nil, "success delete author")
	return resp
}
