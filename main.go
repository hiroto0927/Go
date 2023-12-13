package main

import (
	"myapp/database"
	"myapp/infrastructure/persistence"
	"myapp/interfaces/adapter"
	interfaces "myapp/interfaces/handler"
	"myapp/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	db := database.ConnectDB()
	userRepository := persistence.NewUserPersistence(db)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userHandler := interfaces.NewUserHandler(userUseCase)
	adapter.InitUserAdapter(e, userHandler)

	e.Logger.Fatal(e.Start(":4000"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
