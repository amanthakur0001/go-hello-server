package main

import (
	"github.com/amanthakur0001/go-hello-server/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", api.Hello)
	e.GET("/hello", api.NewHello)
	e.GET("/health", api.Health)
	e.GET("/test", api.Test)
	e.Logger.Fatal(e.Start(":8080"))
}
