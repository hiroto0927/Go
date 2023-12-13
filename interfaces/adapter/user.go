package adapter

import (
	interfaces "myapp/interfaces/handler"

	"github.com/labstack/echo/v4"
)

func InitUserAdapter(e *echo.Echo, uh interfaces.IUserHandler) {
	e.GET("/users/:id", uh.Get())
	e.POST("/users", uh.Post())
}
