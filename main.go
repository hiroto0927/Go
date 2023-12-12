package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/users/:id", getUser)

	e.Logger.Fatal(e.Start(":4000"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type User struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	name := c.QueryParam("name")

	if name == "" {
		name = "anonymous"
	}

	user := User{
		ID:   id,
		Name: &name,
	}

	return c.JSON(http.StatusOK, user)
}
