package main

import (
	"fmt"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/test1", func(c echo.Context) error {
		fmt.Println("TEST1")
		return c.String(http.StatusOK, "TEST1")
	})
	e.GET("/test2", func(c echo.Context) error {
		fmt.Println("TEST2")
		return c.String(http.StatusOK, "TEST1")
	})
	e.GET("/test3", func(c echo.Context) error {
		fmt.Println("TEST3")
		return c.String(http.StatusOK, "TEST1")
	})
	e.GET("/test4", func(c echo.Context) error {
		fmt.Println("TEST4")
		return c.String(http.StatusOK, "TEST1")
	})
	p := prometheus.NewPrometheus("echo", func(c echo.Context) bool {
		return c.Request().RequestURI == "/metrics"
	})
	p.Use(e)
	e.Logger.Fatal(e.Start(":80"))
}
