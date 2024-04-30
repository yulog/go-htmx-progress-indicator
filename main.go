package main

import (
	"fmt"
	"time"

	"github.com/yulog/go-htmx-progress-indicator/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var job = make(chan string)

func main() {
	s := server.NewServer()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.CheckHandler)
	e.GET("/job", s.JobHandler)
	e.GET("/job/progress", s.ProgressHandler)

	e.POST("/start", server.MakeHandler(server.StartHandler, job))

	go func() {
		for j := range job {
			for i := 0; i < 10; i++ {
				s.Progress.Lock()
				s.Progress.Progress += 10
				fmt.Println(j, s.Progress.Progress)
				s.Progress.Unlock()
				time.Sleep(time.Second)
			}
		}
	}()

	e.Logger.Fatal(e.Start(":1324"))
}
