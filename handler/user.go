package handler

import (
	"my-go-starter/model"
	userPkg "my-go-starter/internal/user"
	"my-go-starter/repository"
	"my-go-starter/shared"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	var users []model.User
	users, err := repository.FindAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func MaskUserHandler(c echo.Context) error {
	var user shared.UserType
	if err := c.Bind(&user); err != nil {
		return err
	}
	userMasked := userPkg.MaskUser(user)
	return c.JSON(http.StatusOK, userMasked)

}