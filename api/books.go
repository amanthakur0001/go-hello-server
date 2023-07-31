package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func GetBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, Book{Name: "The Physics Textbook", Author: "RD Sharma", Publisher: "Penguin"})
}
