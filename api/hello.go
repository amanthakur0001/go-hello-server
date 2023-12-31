package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"msg": "hello"})
}

func NewHello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"msg": "This is updated hello!!"})
}
