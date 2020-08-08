package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/previousdeveloper/flagger-dashboard/pkg/controller"
)

type Api struct {
	e *echo.Echo
	c *controller.CanaryController
}

func NewServer(controller *controller.CanaryController) *Api {
	return &Api{e: echo.New(), c: controller}
}

func (s *Api) Start() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
	}))
	s.e.GET("/canary/:namespace", s.c.Hello)
	s.e.Logger.Fatal(s.e.Start(":8888"))
}
