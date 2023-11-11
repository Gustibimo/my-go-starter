package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func HTTPErrorHandler(err error, c echo.Context) {
	type response struct {
		Message string `json:"message"`
	}

	res := response{
		Message: err.Error(),
	}

	c.JSON(http.StatusInternalServerError, res)
}