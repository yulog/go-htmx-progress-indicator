package server

import (
	"sync"

	"github.com/a-h/templ"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Progress *Progress
}

type Progress struct {
	sync.RWMutex
	Progress int
}

func NewServer() *Server {
	return &Server{
		Progress: &Progress{},
	}
}

func MakeHandler(fn func(c echo.Context, ch chan string) error, ch chan string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return fn(c, ch)
	}
}

func renderer(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
