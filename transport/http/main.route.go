package httpRoutes

import (
	"database/sql"
	repositoryMySql "altech-omega-api/repository/mysql"
	handler "altech-omega-api/transport/http/handlers"
	"altech-omega-api/transport/http/middleware"
	"altech-omega-api/usecase"
	"net/http"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger" 
	_ "altech-omega-api/docs" 
)


type Home struct {
	Message string `json:"message"`
}


func homeHandler(c echo.Context) error {
	data := Home{
		Message: "welcome to altech-omega backend testing - Muh. Anshar Azhari",
	}
	return c.JSON(http.StatusOK, data)
}

func StartHttp(e *echo.Echo, db *sql.DB) {
	middleware := middleware.InitMiddleware()
	e.Use(middleware.CORS)

	e.GET("/", homeHandler)

    e.GET("/swagger/", func(c echo.Context) error {
        return c.Redirect(http.StatusMovedPermanently, "/swagger")
    })
    e.GET("/swagger/*", echoSwagger.WrapHandler)

	authorRepo := repositoryMySql.CreateAuthorRepo(db)
	authorUseCase := usecase.CreateAuthorUseCase(authorRepo)
	handler.AuthorRoute(e, authorUseCase)
	
	bookRepo := repositoryMySql.CreateBookRepo(db)
	bookUseCase := usecase.CreateBookUseCase(bookRepo)
	handler.BookRoute(e, bookUseCase)


}
