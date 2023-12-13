package interfaces

import (
	"myapp/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	Get() echo.HandlerFunc
	Post() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.IUserUsecase
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserHandler(uu usecase.IUserUsecase) IUserHandler {
	return &userHandler{
		userUsecase: uu,
	}
}

// Get implements IUserHandler.
func (uh *userHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := uh.userUsecase.GetByUserID(strconv.Itoa(id))

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		userResponse := UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		return c.JSON(http.StatusOK, userResponse)

	}
}

// Post implements IUserHandler.
func (uh *userHandler) Post() echo.HandlerFunc {

	return func(c echo.Context) error {
		userRequest := new(UserRequest)

		if err := c.Bind(userRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err := uh.userUsecase.Insert(userRequest.Name, userRequest.Email)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, "success")
	}

}
