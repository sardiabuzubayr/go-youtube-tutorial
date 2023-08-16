package main

import "github.com/labstack/echo/v4"

func main() {
	route := echo.New()

	route.Start(":8080")
}
