package main

import (
	"log"
	"os"

	"github.com/alexralbino/fasttracktest/quiz-api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339_nano} ${remote_ip} ${host} ${method} ${uri} ${user_agent} ` +
			`${status} ${error} ${latency_human}` + "\n",
	}))
	l := log.New(os.Stdout, "quiz-api", log.LstdFlags)

	qh := handlers.NewQuiz(l)

	// Handle GET Requests
	e.GET("/quiz", qh.StartQuiz)

	// Handle POST Requests
	e.POST("/quiz", qh.AnswerQuiz)

	e.Logger.Fatal(e.Start("localhost:9090"))
}
