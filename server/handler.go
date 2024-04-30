package server

import (
	"fmt"

	cm "github.com/yulog/go-htmx-progress-indicator/components"

	"github.com/labstack/echo/v4"
)

// RootHandler は / のハンドラ
func RootHandler(c echo.Context) error {
	return renderer(c, cm.Index())
}

func (s *Server) CheckHandler(c echo.Context) error {
	s.Progress.RLock()
	defer s.Progress.RUnlock()

	if s.Progress.Progress > 0 {
		return renderer(c, cm.Start())
	}

	return renderer(c, cm.Index())
}

func StartHandler(c echo.Context, ch chan string) error {
	id := c.FormValue("id")

	fmt.Println(id)
	// ジョブキューとかに処理を渡す
	ch <- id

	return renderer(c, cm.Start())
}

func (s *Server) ProgressHandler(c echo.Context) error {
	s.Progress.RLock()
	defer s.Progress.RUnlock()

	if s.Progress.Progress >= 100 {
		c.Response().Header().Set("hx-trigger", "done")
	}

	return renderer(c, cm.Progress(s.Progress.Progress))
}

func (s *Server) JobHandler(c echo.Context) error {
	s.Progress.Lock()
	defer s.Progress.Unlock()
	result := s.Progress.Progress
	s.Progress.Progress = 0

	return renderer(c, cm.Job(result))
}
