package main

import (
	controller "github.com/go-generalize/api_gen/samples/empty_root/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.NewControllers(nil, e)

	// Start echo server
}
