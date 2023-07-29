package main

import (
	"github.com/amanthakur0001/go-hello-server/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", api.Hello)
	e.Logger.Fatal(e.Start(":8080"))
}
